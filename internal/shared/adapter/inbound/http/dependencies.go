package http

import "github.com/labstack/echo/v4"

type Dependencies struct {
	Echo   *echo.Echo
	Prefix string
	Debug  bool
}
