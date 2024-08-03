package restful

import (
	"ninhtq/go-gin/internal/controllers"
)

func (server *Server) enableUserFeatures() {
	router := server.router

	router.Static("/assets", "./public")

	prefixRouter := router.Group("api")

	userController := controllers.NewUserController(server.service.User(), server.config)
	prefixRouter.GET("/users/:id", userController.ReadUser)
	prefixRouter.GET("/users", userController.ReadUsers)
	prefixRouter.POST("/users", userController.CreateUser)
	prefixRouter.PUT("/users", userController.UpdateUser)
	prefixRouter.DELETE("/users", userController.DeleteUser)

	// userRouter := prefixRouter.Group("user")

	/*------------------------ AUTHENTICATED USER ------------------------------*/
	// priRouter := prefixRouter.Group("")
	// priRouter.Use(server.VerifyUserAuthMiddleware())

}
