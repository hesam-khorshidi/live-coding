package inbound

import (
	"context"
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
)

type UserService interface {
	Ingest(ctx context.Context) error
	Get(ctx context.Context, id uuid.UUID) (*domain.User, error)
}
