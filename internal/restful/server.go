package restful

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/services"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config  config.Config
	router  *gin.Engine
	service services.Service
	wg      sync.WaitGroup
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		_ = v.RegisterValidation("enum", func(fl validator.FieldLevel) bool {
			enumString := fl.Param()
			value := fl.Field().String()
			enumSlice := strings.Split(enumString, ";")
			for _, v := range enumSlice {
				if fl.Field().IsZero() || value == v {
					return true
				}
			}
			return false
		})
	}

	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control"}
	router.Use(gin.Logger(), gin.Recovery(), cors.New(corsConf))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "no method provided"})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ready"})
	})

	server.router = router
	server.enableUserFeatures()

	if server.config.Environment != "production" {
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

func (server *Server) Start() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.config.ServerPort),
		Handler: server.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("failed listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Fatalln("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed shutdown server: %s\n", err)
	}

	log.Println("server exiting")

	return nil
}

func (server *Server) Wait() {
	server.wg.Wait()
}

func Init() {
	conf := config.GetConfig()

	server := &Server{
		config: conf,
	}
	server.setupRouter()
	log.Printf("server listen to port %v\n", conf.ServerPort)

	if err := server.Start(); err != nil {
		log.Fatalf("failed to load service %s\n", err)
	}

	server.Wait()
	log.Println("Stopped")
}
