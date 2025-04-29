package user_repo

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	sharedvo "live-coding/internal/shared/domain/value_object"
	"live-coding/internal/user/core/domain"
)

func (r Repository) Get(ctx context.Context, id sharedvo.ID) (*domain.User, error) {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return nil, dbErr
	}

	var user domain.User
	err := db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "user not found")
		}
		return nil, errors.Wrap(err, "db error")
	}

	return &user, nil
}
