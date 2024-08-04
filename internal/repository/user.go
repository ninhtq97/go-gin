package repository

import (
	"errors"
	"fmt"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/core/entities"
	"ninhtq/go-gin/internal/ports"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db DB
}

func NewUserRepository(db DB) ports.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(params ports.CreateUserParams) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	user := entities.NewUser(entities.User{
		Username: params.Username,
		FullName: params.FullName,
		Password: string(hashed),
		Email:    params.Email,
	})

	req := r.db.Client().First(&user, "username = ?", params.Username)
	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	req = r.db.Client().Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	return user.ToDomain(), nil
}

func (r *userRepo) FindMany() ([]*domain.User, error) {
	var users []*domain.User

	req := r.db.Client().Find(&users)
	if req.Error != nil {
		return nil, fmt.Errorf("users not found: %v", req.Error)
	}

	return users, nil
}

func (r *userRepo) FindOne(id string) (*domain.User, error) {
	user := &domain.User{}
	cachekey := user.ID
	err := r.db.cache.Get(cachekey, &user)
	if err == nil {
		return user, nil
	}

	req := r.db.Client().First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err = r.db.cache.Set(cachekey, user, time.Minute*10)
	if err != nil {
		fmt.Printf("Error storing user in cache: %v", err)
	}
	return user, nil
}

func (r *userRepo) Update(id string, params ports.UpdateUserParams) error {
	user := &entities.User{}
	req := r.db.Client().First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	if params.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*params.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("password not hashed: %v", err)
		}
		user.Password = string(hashed)
	}

	if params.FullName != nil {
		user.FullName = *params.FullName
	}

	if params.Email != nil {
		user.Email = *params.Email
	}

	req = r.db.Client().Model(&user).Where("id = ?", id).Updates(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}

	// delete user in the cache
	err := r.db.cache.Delete(string(id))
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}

	return nil
}

func (r *userRepo) Delete(id string) error {
	user := &domain.User{}
	req := r.db.Client().Where("id = ?", id).Delete(&user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}
	err := r.db.cache.Delete(string(id))
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}
	return nil
}
