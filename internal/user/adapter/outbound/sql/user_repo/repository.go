package user_repo

import (
	"live-coding/infra"
	"live-coding/internal/user/core/port/outbound"
)

var _ outbound.UserRepository = (*Repository)(nil)

type Repository struct {
	db *infra.TxDB
}

func New(db *infra.TxDB) Repository {
	return Repository{db: db}
}
