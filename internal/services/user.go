package services

import (
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/internal/ports"
)

type userService struct {
	serviceProperty
}

func NewUserService(property serviceProperty) ports.UserService {
	return &userService{
		serviceProperty: property,
	}
}

func (u *userService) CreateUser(ports.CreateUserInput) (*domain.User, error) {
	panic("unimplemented")
}

func (u *userService) DeleteUser(id uint) error {
	panic("unimplemented")
}

func (u *userService) LoginUser(email string, password string) (*domain.LoginResponse, error) {
	panic("unimplemented")
}

func (u *userService) ReadUser(id uint) (*domain.User, error) {
	panic("unimplemented")
}

func (u *userService) ReadUsers() ([]*domain.User, error) {
	panic("unimplemented")
}

func (u *userService) UpdateUser(id uint, params ports.UpdateUserInput) error {
	panic("unimplemented")
}

func (u *userService) Login(email string, password string) (*domain.LoginResponse, error) {
	panic("unimplemented")
}
