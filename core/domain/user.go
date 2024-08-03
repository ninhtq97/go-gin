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

type LoginResponse struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
