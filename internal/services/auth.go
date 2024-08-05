package services

import (
	"errors"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/internal/ports"
	"ninhtq/go-gin/internal/utils/token"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	serviceProperty
	tokenMarker token.Maker
}

func NewAuthService(property serviceProperty, tokenMarker token.Maker) ports.AuthService {
	return &authService{
		serviceProperty: property,
		tokenMarker:     tokenMarker,
	}
}

func (u *authService) Login(params ports.LoginArgs) (*domain.LoginResponse, error) {
	user, err := u.repo.User().FindOne(ports.FindArgs{
		Username: &params.Username,
	})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, errors.New("unauthorized")
	}

	accessToken, _, err := u.serviceProperty.tokenMaker.CreateToken(user.ID, 24*time.Hour)
	if err != nil {
		return nil, errors.New("can't make access token")
	}

	refreshToken, _, err := u.serviceProperty.tokenMaker.CreateToken(user.ID, 7*24*time.Hour)
	if err != nil {
		return nil, errors.New("can't make refresh token")
	}

	return &domain.LoginResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
