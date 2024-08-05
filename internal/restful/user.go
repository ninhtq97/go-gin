package restful

import (
	"ninhtq/go-gin/internal/controllers"
)

func (server *Server) enableUserFeatures() {
	router := server.router

	router.Static("/assets", "./public")

	prefixRouter := router.Group("api")
	userRouter := prefixRouter.Group("users")

	controller := controllers.NewUserController(server.service.User(), server.config)
	userRouter.POST("", controller.CreateUser)
	userRouter.GET("", controller.ReadUsers)
	userRouter.GET(":id", controller.ReadUser)
	userRouter.PATCH("/:id", controller.UpdateUser)
	userRouter.DELETE("/:id", controller.DeleteUser)

	/*------------------------ AUTHENTICATED USER ------------------------------*/
	// priRouter := prefixRouter.Group("")
	// priRouter.Use(server.VerifyUserAuthMiddleware())

}
