package user_repo

import (
	"context"
	"live-coding/internal/user/adapter/outbound/sql/model"
	"live-coding/internal/user/core/domain"
)

func (r Repository) Save(ctx context.Context, user domain.User) error {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return dbErr
	}

	um := model.UserToModel(user)
	_, err := db.NewInsert().
		Model(&um).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
