package entities

import (
	"ninhtq/go-gin/core/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Code     string `gorm:"code"`
	FullName string `gorm:"fullName"`
	Email    string `gorm:"email"`
}

func (data User) ToDomain() domain.User {
	return domain.User{
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
		Code:     arg.Code,
		FullName: arg.FullName,
	}
}
