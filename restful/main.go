package restful

import (
	"net/http"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/services"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Config  config.Config
	Router  *gin.Engine
	Service services.Service
	Wg      sync.WaitGroup
}

func (server *Server) SetupRouter() {
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

	cnf := cors.DefaultConfig()
	cnf.AllowAllOrigins = true
	cnf.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control"}
	router.Use(cors.New(cnf))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	})
	router.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "no method provided"})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ready"})
	})

	server.Router = router
	server.enableFeatures()
}

func (server *Server) enableFeatures() {
	server.enableUserFeatures()

	if server.Config.Environment != "production" {
		server.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
