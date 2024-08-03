package main

import (
	"log"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/core/entities"
	"ninhtq/go-gin/internal/cache"
	"ninhtq/go-gin/internal/repository"
	"ninhtq/go-gin/internal/restful"
	"ninhtq/go-gin/internal/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userService *services.UserService
)

// @Title 			Bvote API
// @Version         1.0
// @Description 	This is a server Bvote
// @Schemes			http https

// @BasePath		/api

// @SecurityDefinitions.apikey Bearer
// @In header
// @Name authorization
// @Description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	config.Init(".env")
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

	store := repository.NewDB(db, redisCache)

	userService = services.NewUserService(store)

	restful.Init()
}
