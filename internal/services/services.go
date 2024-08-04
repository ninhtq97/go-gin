package services

import (
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"
)

type serviceProperty struct {
	config config.Config
	repo   ports.Repository
}

type services struct {
	property    serviceProperty
	userService ports.UserService
}

func NewService(config config.Config, repo ports.Repository) (ports.Service, error) {
	property := serviceProperty{
		config: config,
		repo:   repo,
	}

	svc := services{
		property:    property,
		userService: NewUserService(property),
	}

	return &svc, nil
}

func (s *services) User() ports.UserService {
	return s.userService
}
