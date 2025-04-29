package user_srv

import (
	"live-coding/config"
	"live-coding/internal/user/core/port/inbound"
	"live-coding/internal/user/core/port/outbound"
)

var _ inbound.UserService = (*Service)(nil)

type Service struct {
	userRepo    outbound.UserRepository
	userFileSrv outbound.UserReaderService
	workerCount int
}

func New(
	userRepo outbound.UserRepository,
	userFileSrv outbound.UserReaderService,
	workerConfig config.WorkerConfig,
) Service {
	return Service{
		userRepo:    userRepo,
		userFileSrv: userFileSrv,
		workerCount: workerConfig.UserWorkerCount,
	}
}
