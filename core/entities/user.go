package entities

import (
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/internal/utils"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Entity
	Username string `gorm:"notnull;unique"`
	Password string `gorm:"notnull"`
	Code     string `gorm:"notnull;unique"`
	FullName string `gorm:"notnull"`
	Email    string `gorm:"notnull;unique"`
}

func (data User) ToDomain() *domain.User {
	return &domain.User{
		ID:        data.Entity.ID,
		CreatedAt: data.Entity.CreatedAt,
		Code:      data.Code,
		FullName:  data.FullName,
		Email:     data.Email,
	}
}

func AsUser(arg domain.User) User {
	return User{
		Entity: Entity{
			ID:        arg.ID,
			CreatedAt: arg.CreatedAt,
		},
		Username: arg.Username,
		Code:     arg.Code,
		FullName: arg.FullName,
		Email:    arg.Email,
	}
}

func NewUser(arg User) *User {
	now := time.Now()
	code := utils.MakeStr(12, "")

	return &User{
		Entity: Entity{
			ID:        uuid.New().String(),
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
