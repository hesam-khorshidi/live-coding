package infra

import (
	"context"
	"errors"
	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"slices"
)

type HTTPServerConfig struct {
	Debug      bool
	Protocol   string
	Host       string
	Port       string
	ApiPrefix  string
	ApiVersion string
	LogEnable  bool
}

// NewHttpServer godoc
// @title           Live Coding Session
// @version         1.0
// @description     Live Coding
// @host      localhost:8080
// @BasePath  /api/v1
func NewHttpServer(txDB *TxDB) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(WithTransaction(txDB))

	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("5M"))

	e.GET("/reference", func(c echo.Context) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Live Coding Document",
			},
			DarkMode: true,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, "Error generating API reference")
		}

		return c.HTML(http.StatusOK, htmlContent)
	})

	return e, nil
}

var failedStatuses = []int{http.StatusBadRequest, http.StatusInternalServerError}

func WithTransaction(txDB *TxDB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			txerr := txDB.RunInTransaction(c.Request().Context(), func(ctx context.Context) bool {
				c.SetRequest(c.Request().WithContext(ctx))
				err = next(c)
				if err != nil {
					return false
				}

				if slices.Contains(failedStatuses, c.Response().Status) {
					return false
				}

				return true
			})
			return errors.Join(err, txerr)
		}
	}
}
