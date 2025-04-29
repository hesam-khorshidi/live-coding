package infra

import (
	"context"
	"errors"
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

func NewHttpServer(txDB *TxDB) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(WithTransaction(txDB))

	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("5M"))
	e.Static("/", "public")

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
