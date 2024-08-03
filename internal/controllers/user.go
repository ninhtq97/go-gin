package controllers

import (
	"net/http"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc    ports.UserService
	config config.Config
}

func NewUserController(UserService ports.UserService, config config.Config) *UserController {
	return &UserController{
		svc:    UserService,
		config: config,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user ports.CreateUserInput
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err := c.svc.CreateUser(user)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New user created successfully",
	})
}

func (c *UserController) ReadUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.svc.ReadUser(uint(id))

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) ReadUsers(ctx *gin.Context) {
	users, err := c.svc.ReadUsers()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	// Load API configuration
	// apiCfg, err := repository.LoadAPIConfig()
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// // Validate token
	// userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// // Update user
	// var user domain.User
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// err = h.svc.UpdateUser(userID, user.Email, user.Password)
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	// apiCfg, err := repository.LoadAPIConfig()
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// userID, err := ValidateToken(ctx.Request.Header.Get("Authorization"), apiCfg.JWTSecret)
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	// err = h.svc.DeleteUser(userID)
	// if err != nil {
	// 	HandleError(ctx, http.StatusBadRequest, err)
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
