package outbound_prt

import (
	"context"
	sharedvo "live-coding/internal/shared/domain/value_object"
	"live-coding/internal/user/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	Get(ctx context.Context, id sharedvo.ID) (*domain.User, error)
}
