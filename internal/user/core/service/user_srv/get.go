package user_srv

import (
	"context"
	sharedvo "live-coding/internal/shared/domain/value_object"
	"live-coding/internal/user/core/domain"
)

func (s Service) Get(ctx context.Context, id sharedvo.ID) (*domain.User, error) {
	return s.userRepo.Get(ctx, id)
}
