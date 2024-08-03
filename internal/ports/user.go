package ports

import "ninhtq/go-gin/core/domain"

type CreateUserInput struct {
	Username string
	Password string
	FullName string
	Email    string
}

type UpdateUserInput struct {
	Password *string
	FullName *string
	Email    *string
}

type UserRepository interface {
	CreateUser(CreateUserInput) (*domain.User, error)
	ReadUser(id uint) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(id uint, input UpdateUserInput) error
	DeleteUser(id uint) error
	LoginUser(email, password string) (*domain.LoginResponse, error)
}
