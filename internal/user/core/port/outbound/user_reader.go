package outbound

import (
	"context"
	"live-coding/internal/user/core/domain"
)

type UserReaderService interface {
	ReadUsers(ctx context.Context) ([]domain.User, error)
}
