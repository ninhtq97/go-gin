package repository

import (
	"errors"
	"fmt"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/core/entities"
	"ninhtq/go-gin/core/exception"
	"ninhtq/go-gin/internal/ports"
	"ninhtq/go-gin/internal/utils/token"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (u *DB) CreateUser(input ports.CreateUserInput) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password not hashed: %v", err)
	}

	user := entities.NewUser(entities.User{
		Username: input.Username,
		FullName: input.FullName,
		Password: string(hashed),
		Email:    input.Email,
	})

	req := u.db.First(&user, "email = ?", input.Email)
	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	req = u.db.Create(&user)
	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}
	return user.ToDomain(), nil
}

func (u *DB) ReadUser(id uint) (*domain.User, error) {
	user := &domain.User{}
	cachekey := user.ID
	err := u.cache.Get(string(cachekey), &user)
	if err == nil {
		return user, nil
	}

	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err = u.cache.Set(string(cachekey), user, time.Minute*10)
	if err != nil {
		fmt.Printf("Error storing user in cache: %v", err)
	}
	return user, nil
}

func (u *DB) ReadUsers() ([]*domain.User, error) {
	var users []*domain.User

	req := u.db.Find(&users)
	if req.Error != nil {
		return nil, fmt.Errorf("users not found: %v", req.Error)
	}

	return users, nil
}

func (u *DB) UpdateUser(id uint, input ports.UpdateUserInput) error {
	user := &entities.User{}
	req := u.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}

	if input.Password != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("password not hashed: %v", err)
		}
		user.Password = string(hashed)
	}

	if input.FullName != nil {
		user.Email = *input.FullName
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	user.Email = *input.Email

	req = u.db.Model(&user).Where("id = ?", id).Updates(user)
	if req.RowsAffected == 0 {
		return errors.New("unable to update user :(")
	}

	// delete user in the cache
	err := u.cache.Delete(string(id))
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}

	return nil
}

func (u *DB) DeleteUser(id uint) error {
	user := &domain.User{}
	req := u.db.Where("id = ?", id).Delete(&user)
	if req.RowsAffected == 0 {
		return errors.New("user not found")
	}
	err := u.cache.Delete(string(id))
	if err != nil {
		fmt.Printf("Error deleting user in cache: %v", err)
	}
	return nil
}

func (u *DB) LoginUser(email, password string) (*domain.LoginResponse, error) {
	conf := config.GetConfig()

	user, err := u.findUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = u.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.generateAccessToken(user.ID, conf.JWTSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.generateRefreshToken(user.ID, conf.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *DB) findUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	req := u.db.First(&user, "email = ?", email)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *DB) VerifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password not matched")
	}
	return nil
}

func (u *DB) generateAccessToken(userID uint, jwtSecret string) (string, error) {
	payload, err := token.NewPayload(userID, 24*time.Hour)
	if err != nil {
		return "", exception.New(exception.TypeValidation, "Sign token failed. Pls try again!", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(jwtSecret))
}

func (u *DB) generateRefreshToken(userID uint, jwtSecret string) (string, error) {
	payload, err := token.NewPayload(userID, 7*24*time.Hour)
	if err != nil {
		return "", exception.New(exception.TypeValidation, "Sign token failed. Pls try again!", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(jwtSecret))
}
