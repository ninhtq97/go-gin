package repository

import (
	"ninhtq/go-gin/internal/cache"

	"gorm.io/gorm"
)

type DB struct {
	db    *gorm.DB
	cache *cache.RedisCache
}

func NewDB(db *gorm.DB, cache *cache.RedisCache) *DB {
	return &DB{
		db:    db,
		cache: cache,
	}
}
