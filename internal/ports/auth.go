package ports

import (
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/internal/utils/token"
)

type LoginArgs struct {
	Username string
	Password string
}

type AuthService interface {
	Login(LoginArgs) (*domain.LoginResponse, error)
}

type AuthDecoded struct {
	Token   string
	Payload *token.Payload
}
