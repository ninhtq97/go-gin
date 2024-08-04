package repository

import (
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"
)

type repos struct {
	db       DB
	rdb      DB
	config   config.Config
	userRepo ports.UserRepository
}

func NewRepository(db DB, config config.Config) ports.Repository {
	return &repos{
		db:       db,
		userRepo: NewUserRepository(db),
		config:   config,
	}
}

func (r *repos) User() ports.UserRepository {
	return r.userRepo
}
