package services

import (
	"context"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/utils/token"
)

type AuthParams struct {
	Token   string
	Payload *token.Payload
}

type GetUserParams struct {
	ID    uint
	Email string
}

type CreateUserParams struct {
	User domain.User
}

type UpdateUserParams struct {
	User domain.User
}

type DeleteUserParams struct {
	ID string
}

type LoginParams struct {
	User domain.User
}

type UserService interface {
	Get(context.Context, GetUserParams) (domain.User, error)
	Create(context.Context, CreateUserParams) (domain.User, error)
	Update(context.Context, UpdateUserParams) (domain.User, error)
	Delete(context.Context, DeleteUserParams) error
	Login(context.Context, LoginParams) (string, error)
}
