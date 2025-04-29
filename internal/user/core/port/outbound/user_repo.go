package outbound

import (
	"context"
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	Get(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
