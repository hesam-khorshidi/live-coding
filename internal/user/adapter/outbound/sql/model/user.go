package model

import (
	"github.com/uptrace/bun"
	"live-coding/internal/user/core/domain"
	"live-coding/pkg/slice"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64     `bun:"pk,id"`
	Name          string    `bun:"name"`
	Email         string    `bun:"email"`
	PhoneNumber   string    `bun:"phone_number"`
	Addresses     []Address `bun:"rel:has-many,join:user_id"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Addresses:   slice.Convert(u.Addresses, AddressToDomain),
	}
}

func UserToModel(user domain.User) User {
	addresses := slice.Convert(user.Addresses, AddressToModel)
	for _, address := range addresses {
		address.UserID = int64(user.ID)
	}

	return User{
		ID:          int64(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Addresses:   addresses,
	}
}
