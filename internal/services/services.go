package services

import (
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"
	"ninhtq/go-gin/internal/utils/token"
)

type serviceProperty struct {
	config     config.Config
	repo       ports.Repository
	tokenMaker token.Maker
}

type services struct {
	property    serviceProperty
	authService ports.AuthService
	userService ports.UserService
}

func NewService(config config.Config, repo ports.Repository) (ports.Service, error) {
	tokenMaker, err := token.NewJWTMaker(config.JWTSecret)
	if err != nil {
		return nil, err
	}

	property := serviceProperty{
		config:     config,
		repo:       repo,
		tokenMaker: tokenMaker,
	}

	svc := services{
		property:    property,
		authService: NewAuthService(property),
		userService: NewUserService(property),
	}

	return &svc, nil
}

func (s *services) TokenMaker() token.Maker {
	return s.property.tokenMaker
}

func (s *services) Auth() ports.AuthService {
	return s.authService
}

func (s *services) User() ports.UserService {
	return s.userService
}
