package ports

import (
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/internal/repository"
)

type UserRepository interface {
	CreateUser(email, password string) (*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(id, email, password string) error
	DeleteUser(id string) error
	LoginUser(email, password string) (*repository.LoginResponse, error)
	UpdateMembershipStatus(id string, status bool) error
}
