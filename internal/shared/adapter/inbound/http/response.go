package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty" swaggerignore:"true"`
	Message string `json:"message,omitempty" swaggerignore:"true"`
	Meta    any    `json:"meta,omitempty" swaggerignore:"true"`
}

type ListMeta struct {
	TotalCount int `json:"total_count"`
}

func InternalError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: err.Error(),
	})
}

func BadRequest(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: err.Error(),
	})
}

func PaymentRequired(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusPaymentRequired, Response{
		Success: false,
		Message: err.Error(),
	})
}

func Notfound(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: err.Error(),
	})
}

func Success(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func SuccessWithMeta(ctx echo.Context, data, meta any) error {
	return ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func SuccessWithListMeta(ctx echo.Context, data any, totalCount int) error {
	return SuccessWithMeta(ctx, data, ListMeta{TotalCount: totalCount})
}

func NoContent(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func Forbidden(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusForbidden, Response{Success: false, Message: "access denied!"})
}

func TooManyRequest(ctx echo.Context) error {
	return ctx.NoContent(http.StatusTooManyRequests)
}

func Unauthorized(ctx echo.Context, _ error) error {
	return ctx.JSON(http.StatusUnauthorized, Response{Success: false, Message: "unauthorized!"})
}
