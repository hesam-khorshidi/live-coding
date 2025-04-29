package domain

import (
	"github.com/google/uuid"
)

type Address struct {
	ID      uuid.UUID
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}
