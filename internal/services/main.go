package services

import "ninhtq/go-gin/internal/utils/token"

type Service interface {
	TokenMaker() token.Maker
	User() UserService
}
