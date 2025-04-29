package dto

import (
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
)

type Address struct {
	ID      uuid.UUID `json:"id"`
	Street  string    `json:"street"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
}

func AddressToDomain(u Address) domain.Address {
	return domain.Address{
		ID:      u.ID,
		Street:  u.Street,
		City:    u.City,
		State:   u.State,
		ZipCode: u.ZipCode,
		Country: u.Country,
	}
}
