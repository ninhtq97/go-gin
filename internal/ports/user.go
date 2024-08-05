package ports

import (
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/core/entities"
)

type CreateUserArgs struct {
	Username string
	Password string
	FullName string
	Email    string
}

type UpdateUserArgs struct {
	Password *string
	FullName *string
	Email    *string
}

type UserService interface {
	CreateUser(args CreateUserArgs) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	ReadUser(id string) (*domain.User, error)
	UpdateUser(id string, args UpdateUserArgs) error
	DeleteUser(id string) error
}

type FindArgs struct {
	ID       *string
	Code     *string
	Username *string
	FullName *string
	Email    *string
}

type UserRepository interface {
	Create(args CreateUserArgs) (*domain.User, error)
	FindMany() ([]*domain.User, error)
	FindOne(FindArgs) (*entities.User, error)
	Update(id string, args UpdateUserArgs) error
	Delete(id string) error
}
