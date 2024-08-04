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

func (u *userService) CreateUser(input ports.CreateUserParams) (*domain.User, error) {
	entity, err := u.repo.User().Create(input)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *userService) ReadUsers() ([]*domain.User, error) {
	entities, err := u.repo.User().FindMany()
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (u *userService) ReadUser(id string) (*domain.User, error) {
	entity, err := u.repo.User().FindOne(id)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *userService) UpdateUser(id string, params ports.UpdateUserParams) error {
	err := u.repo.User().Update(id, params)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) DeleteUser(id string) error {
	err := u.repo.User().Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Login(email string, password string) (*domain.LoginResponse, error) {
	panic("unimplemented")
}
