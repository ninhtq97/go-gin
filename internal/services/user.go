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

func (u *userService) CreateUser(args ports.CreateUserArgs) (*domain.User, error) {
	entity, err := u.repo.User().Create(args)

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
	entity, err := u.repo.User().FindOne(ports.FindArgs{
		ID: &id,
	})
	if err != nil {
		return nil, err
	}

	return entity.ToDomain(), nil
}

func (u *userService) UpdateUser(id string, args ports.UpdateUserArgs) error {
	err := u.repo.User().Update(id, args)
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
