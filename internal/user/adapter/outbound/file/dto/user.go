package dto

import (
	"github.com/google/uuid"
	"live-coding/internal/user/core/domain"
	"live-coding/pkg/slice"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Addresses   []Address `json:"addresses"`
}

func UserToDomain(u User) domain.User {
	return domain.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Addresses:   slice.Convert(u.Addresses, AddressToDomain),
	}
}
