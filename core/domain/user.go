package domain

import (
	"time"
)

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Code      string
	FullName  string
	Email     string
}
