package db

import (
	"log"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/core/entities"
	"ninhtq/go-gin/internal/cache"
	"ninhtq/go-gin/internal/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userService *services.UserService
)

type DB struct {
	db    *gorm.DB
	cache *cache.RedisCache
}

func Init() {
	conf := config.GetConfig()

	db, err := gorm.Open(mysql.Open(conf.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s\n", err)
	}

	redisCache, err := cache.NewRedisCache(conf.RedisSource, "")
	if err != nil {
		log.Fatalf("failed to connect redis: %s\n", err)
	}

	db.AutoMigrate(&entities.User{})

	store := NewDB(db, redisCache)

	userService = services.NewUserService(store)
}

func NewDB(db *gorm.DB, cache *cache.RedisCache) *DB {
	return &DB{
		db:    db,
		cache: cache,
	}
}
