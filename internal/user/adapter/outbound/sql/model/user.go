package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"live-coding/internal/user/core/domain"
	"live-coding/pkg/slice"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `bun:"id,pk,type:uuid"`
	Name          string    `bun:"name"`
	Email         string    `bun:"email"`
	PhoneNumber   string    `bun:"phone_number"`
	Addresses     []Address `bun:"rel:has-many,join:id=user_id"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Addresses:   slice.Convert(u.Addresses, AddressToDomain),
	}
}

func UserToModel(user domain.User) User {
	addresses := slice.Convert(user.Addresses, AddressToModel)
	for i := range addresses {
		addresses[i].UserID = user.ID
	}

	return User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Addresses:   addresses,
	}
}
