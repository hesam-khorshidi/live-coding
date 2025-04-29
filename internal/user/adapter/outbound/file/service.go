package file

import (
	"live-coding/config"
	"live-coding/internal/user/core/port/outbound"
)

var _ outbound.UserReaderService = (*Service)(nil)

type Service struct {
	FilePath string
}

func New(cfg config.FileConfig) Service {
	return Service{FilePath: cfg.UserJson}
}
