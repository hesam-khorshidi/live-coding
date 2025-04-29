package domain

import sharedvo "live-coding/internal/shared/domain/value_object"

type Address struct {
	ID      sharedvo.ID
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}
