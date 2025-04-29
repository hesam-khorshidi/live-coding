package inbound_prt

import (
	"context"
	sharedvo "live-coding/internal/shared/domain/value_object"
	"live-coding/internal/user/core/domain"
)

type UserService interface {
	Ingest(ctx context.Context) error
	Get(ctx context.Context, id sharedvo.ID) (*domain.User, error)
}
