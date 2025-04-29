package user_srv

import (
	inbound "live-coding/internal/user/core/port/inbound_prt"
	"live-coding/internal/user/core/port/outbound_prt"
)

var _ inbound.UserService = (*Service)(nil)

type Service struct {
	userRepo outbound_prt.UserRepository
}

func NewService(userRepo outbound_prt.UserRepository) Service {
	return Service{userRepo: userRepo}
}
