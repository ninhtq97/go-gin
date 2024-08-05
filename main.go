package main

import (
	"log"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/core/entities"
	_ "ninhtq/go-gin/docs"
	"ninhtq/go-gin/internal/cache"
	"ninhtq/go-gin/internal/repository"
	"ninhtq/go-gin/internal/restful"
	"ninhtq/go-gin/internal/services"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @Title 				Go Gin API
// @Version       1.0
// @Description		This is a server for development
// @Schemes				http https

// @BasePath			/api

// @SecurityDefinitions.apikey Bearer
// @In header
// @Name authorization
// @Description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	conf := config.Init(".env")

	db, err := gorm.Open(mysql.Open(conf.DBSource), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      true,          // Don't include params in the SQL log
				Colorful:                  false,         // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %s\n", err)
	}

	redisCache, err := cache.NewRedisCache(conf.RedisSource, "")
	if err != nil {
		log.Fatalf("failed to connect redis: %s\n", err)
	}

	db.AutoMigrate(&entities.User{})

	store := repository.NewDB(db, redisCache)

	repo := repository.NewRepository(*store, conf)
	log.Println("load repository done")

	service, err := services.NewService(conf, repo)
	if err != nil {
		log.Fatalf("failed to load service %s\n", err)
	}

	log.Printf("server listen to port %d", conf.ServerPort)
	server := restful.NewServer(conf, service)

	if err := server.Start(); err != nil {
		log.Fatalf("failed to load service %s\n", err)
	}

	server.Wait()
	log.Println("Stopped")
}
