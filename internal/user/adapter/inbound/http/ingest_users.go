package http

import (
	"github.com/labstack/echo/v4"
	"live-coding/internal/shared/adapter/inbound/http"
)

// IngestUsers godoc
// @Summary      Ingest users
// @Description  Start user ingestion from json file
// @Tags         User
// @Produce      json
// @Success      200  {object}	http.Response
// @Failure      500  {object}  http.Response
// @Router       /users/ingest [post]
func (c Controller) IngestUsers(ctx echo.Context) error {
	err := c.userSrv.Ingest(ctx.Request().Context())
	if err != nil {
		return http.InternalError(ctx, err)
	}

	return http.Success(ctx, nil)
}
