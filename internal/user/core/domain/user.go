package domain

import sharedvo "live-coding/internal/shared/domain/value_object"

type User struct {
	ID          sharedvo.ID
	Name        string
	Email       string
	PhoneNumber string
	Addresses   []Address
}
