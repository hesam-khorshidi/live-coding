package http

import (
	"live-coding/internal/shared/adapter/inbound/http"
	"live-coding/internal/user/core/port/inbound"
)

type Controller struct {
	userSrv inbound.UserService
}

func Init(d http.Dependencies, userSrv inbound.UserService) Controller {
	c := Controller{userSrv}

	g := d.Echo.Group(d.Prefix + "/users")
	g.GET("/:id", c.GetUser)
	g.POST("/ingest", c.IngestUsers)

	return c
}
