package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"live-coding/internal/user/core/domain"
)

type Address struct {
	bun.BaseModel `bun:"table:addresses,alias:ua"`
	ID            uuid.UUID `bun:"id,pk,type:uuid"`
	UserID        uuid.UUID `bun:"user_id,type:uuid"`
	Street        string    `bun:"street"`
	City          string    `bun:"city"`
	State         string    `bun:"state"`
	ZipCode       string    `bun:"zip_code"`
	Country       string    `bun:"country"`
}

func AddressToDomain(a Address) domain.Address {
	return domain.Address{
		ID:      a.ID,
		Street:  a.Street,
		City:    a.City,
		State:   a.State,
		ZipCode: a.ZipCode,
		Country: a.Country,
	}
}

func AddressToModel(a domain.Address) Address {
	return Address{
		ID:      a.ID,
		Street:  a.Street,
		City:    a.City,
		State:   a.State,
		ZipCode: a.ZipCode,
		Country: a.Country,
	}
}
