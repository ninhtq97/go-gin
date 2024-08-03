package entities

import (
	"ninhtq/go-gin/core/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username;notnull"`
	Password string `gorm:"password;notnull"`
	Code     string `gorm:"code;notnull"`
	FullName string `gorm:"fullName;notnull"`
	Email    string `gorm:"email;notnull"`
}

func (data User) ToDomain() *domain.User {
	return &domain.User{
		ID:        data.Model.ID,
		Code:      data.Code,
		FullName:  data.FullName,
		Email:     data.Email,
		CreatedAt: data.Model.CreatedAt,
		UpdatedAt: data.Model.UpdatedAt,
	}
}

func AsUser(arg domain.User) User {
	return User{
		Model: gorm.Model{
			ID:        arg.ID,
			CreatedAt: arg.CreatedAt,
			UpdatedAt: arg.UpdatedAt,
		},
		Username: arg.Username,
		Code:     arg.Code,
		FullName: arg.FullName,
		Email:    arg.Email,
	}
}

func NewUser(arg User) User {
	now := time.Now()
	code := uuid.New().NodeID()

	return User{
		Model: gorm.Model{
			ID:        uint(uuid.New().ID()),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Username: arg.Username,
		Password: arg.Password,
		Code:     string(code),
		FullName: arg.FullName,
		Email:    arg.Email,
	}
}
