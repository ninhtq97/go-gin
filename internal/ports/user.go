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

type UserService interface {
	CreateUser(CreateUserInput) (*domain.User, error)
	ReadUser(id uint) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(id uint, params UpdateUserInput) error
	DeleteUser(id uint) error
	Login(email, password string) (*domain.LoginResponse, error)
}

type UserRepository interface {
	Create(CreateUserInput) (*domain.User, error)
	FindMany() ([]*domain.User, error)
	FindOne(id uint) (*domain.User, error)
	Update(id uint, input UpdateUserInput) error
	Delete(id uint) error
}
