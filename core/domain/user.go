package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	Code      string    `json:"code"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
}

type LoginResponse struct {
	ID           string `json:"id"`
	FullName     string `json:"fullName"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
