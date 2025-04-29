package http

import "github.com/labstack/echo/v4"

type Dependencies struct {
	Echo        *echo.Echo
	Prefix      string
	Middlewares Middlewares
	Debug       bool
}

type Middlewares struct {
	JWT         echo.MiddlewareFunc
	APIKey      echo.MiddlewareFunc
	BindUser    echo.MiddlewareFunc
	Synchronize func(routeKey string, perUser bool) echo.MiddlewareFunc
}
