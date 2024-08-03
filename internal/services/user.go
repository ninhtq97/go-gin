package services

import "ninhtq/go-gin/internal/ports"

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
