package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) enableUserFeatures() {
	router := server.Router

	router.Static("/assets", "./public")

	prefixRouter := router.Group("api")

	prefixRouter.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// userRouter := prefixRouter.Group("user")

	/*------------------------ AUTHENTICATED USER ------------------------------*/
	// priRouter := prefixRouter.Group("")
	// priRouter.Use(server.VerifyUserAuthMiddleware())

}
