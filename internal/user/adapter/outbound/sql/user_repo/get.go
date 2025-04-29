package user_repo

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"live-coding/internal/user/adapter/outbound/sql/model"
	"live-coding/internal/user/core/domain"
)

func (r Repository) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return nil, dbErr
	}

	var user model.User
	err := db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Relation("Addresses").
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "user not found")
		}
		return nil, errors.Wrap(err, "db error")
	}

	ud := user.ToDomain()
	return &ud, nil
}
