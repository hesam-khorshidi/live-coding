package http

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func UUIDParamLoader(ctx echo.Context, key string) (*uuid.UUID, error) {
	str := ctx.Param(key)
	if str == "" {
		str = ctx.QueryParam(key)
	}

	if str == "" {
		return nil, fmt.Errorf("missing parameter: %s", key)
	}

	u, err := uuid.Parse(str)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format for parameter %s: %v", key, err)
	}
	return &u, nil

}
