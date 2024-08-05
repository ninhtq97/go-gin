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

func (r *userRepo) Create(args ports.CreateUserArgs) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	user := entities.NewUser(entities.User{
		Username: args.Username,
		FullName: args.FullName,
		Password: string(hashed),
		Email:    args.Email,
	})

	req := r.db.Client().First(&user, "username = ?", args.Username)
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

func (r *userRepo) FindOne(args ports.FindArgs) (*entities.User, error) {
	user := &entities.User{}
	cachekey := user.ID
	err := r.db.cache.Get(cachekey, &user)
	if err == nil {
		return user, nil
	}

	client := r.db.Client()

	if args.ID != nil {
		client = client.Where("id = ?", *args.ID)
	}

	if args.Code != nil {
		client = client.Where("code = ?", *args.Code)
	}

	if args.Username != nil {
		client = client.Where("username = ?", *args.Username)
	}

	if args.FullName != nil {
		client = client.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", *args.FullName))
	}

	if args.Email != nil {
		client = client.Where("email LIKE ?", fmt.Sprintf("%%%s%%", *args.Email))
	}

	req := client.First(&user)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err = r.db.cache.Set(cachekey, user, time.Minute*10)
	if err != nil {
		fmt.Printf("Error storing user in cache: %v", err)
	}
	return user, nil
}

func (r *userRepo) Update(id string, args ports.UpdateUserArgs) error {
	user := &entities.User{}
	req := r.db.Client().First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	if args.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*args.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("password not hashed: %v", err)
		}
		user.Password = string(hashed)
	}

	if args.FullName != nil {
		user.FullName = *args.FullName
	}

	if args.Email != nil {
		user.Email = *args.Email
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
