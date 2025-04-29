package user_repo

import (
	"live-coding/infra"
	outbound "live-coding/internal/user/core/port/outbound_prt"
)

var _ outbound.UserRepository = (*Repository)(nil)

type Repository struct {
	db *infra.TxDB
}

func NewRepository(db *infra.TxDB) Repository {
	return Repository{db: db}
}
