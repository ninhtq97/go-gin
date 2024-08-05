package ports

import "ninhtq/go-gin/internal/utils/token"

type Server interface {
	Start() error
	Wait()
}

type Service interface {
	TokenMaker() token.Maker
	Auth() AuthService
	User() UserService
}

type Repository interface {
	User() UserRepository
}
