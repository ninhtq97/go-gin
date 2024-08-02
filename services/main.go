package services

import (
	"ninhtq/go-gin/utils/token"
)

type Service interface {
	TokenMaker() token.Maker
	User() UserService
}
