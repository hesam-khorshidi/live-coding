package model

import (
	"github.com/uptrace/bun"
	"live-coding/internal/user/core/domain"
)

type Address struct {
	bun.BaseModel `bun:"table:user_addresses,alias:ua"`
	ID            int64  `bun:"pk,id"`
	UserID        int64  `bun:"user_id"`
	Street        string `bun:"street"`
	City          string `bun:"city"`
	State         string `bun:"state"`
	ZipCode       string `bun:"zip_code"`
	Country       string `bun:"country"`
}

func AddressToDomain(a Address) domain.Address {
	return domain.Address{
		Street:  a.Street,
		City:    a.City,
		State:   a.State,
		ZipCode: a.ZipCode,
		Country: a.Country,
	}
}

func AddressToModel(a domain.Address) Address {
	return Address{
		ID:      int64(a.ID),
		Street:  a.Street,
		City:    a.City,
		State:   a.State,
		ZipCode: a.ZipCode,
		Country: a.Country,
	}
}
