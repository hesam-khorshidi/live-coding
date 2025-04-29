package user_srv

import (
	"context"
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
)

func (s Service) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepo.Get(ctx, id)
}
