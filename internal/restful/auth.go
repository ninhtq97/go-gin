package restful

import (
	"ninhtq/go-gin/internal/controllers"
)

func (server *Server) enableAuthFeatures() {
	router := server.router

	prefixRouter := router.Group("api")

	controller := controllers.NewAuthController(server.service.Auth(), server.config)
	prefixRouter.POST("login", controller.Login)

	// userRouter := prefixRouter.Group("user")

	/*------------------------ AUTHENTICATED USER ------------------------------*/
	priRouter := prefixRouter.Group("")
	priRouter.Use(server.VerifyUserAuthMiddleware())

	priRouter.GET("me", controller.Me)
}
