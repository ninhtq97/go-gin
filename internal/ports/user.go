package ports

import "ninhtq/go-gin/core/domain"

type CreateUserParams struct {
	Username string
	Password string
	FullName string
	Email    string
}

type UpdateUserParams struct {
	Password *string
	FullName *string
	Email    *string
}

type UserService interface {
	CreateUser(params CreateUserParams) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	UpdateUser(id string, params UpdateUserParams) error
	DeleteUser(id string) error
	Login(email, password string) (*domain.LoginResponse, error)
}

type UserRepository interface {
	Create(params CreateUserParams) (*domain.User, error)
	FindMany() ([]*domain.User, error)
	FindOne(id string) (*domain.User, error)
	Update(id string, params UpdateUserParams) error
	Delete(id string) error
}
