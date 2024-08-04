package controllers

import (
	"net/http"
	"net/mail"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"

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

// CreateUser godoc
//
// @Description 	Get access token
// @Tags 					User
// @Accept 				json
// @Produce 			json
// @Param					auth		body				CreateUserRequest					true		"Auth admin"
// @Success 			200 		{object}		object{message=string}
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/users	[POST]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var json CreateUserRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err := c.svc.CreateUser(ports.CreateUserParams{
		Username: json.Username,
		Password: json.Password,
		FullName: json.FullName,
		Email:    json.Email,
	})
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New user created successfully",
	})
}

// ReadUsers godoc
//
// @Description 	Get list users
// @Tags 					User
// @Accept 				json
// @Produce 			json
// @Success 			200 		{array}			domain.User
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/users	[GET]
func (c *UserController) ReadUsers(ctx *gin.Context) {
	users, err := c.svc.ReadUsers()
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// ReadUser godoc
//
// @Description 	Get user
// @Tags 					User
// @Accept 				json
// @Produce 			json
// @Param					id			path				string								true		"User ID"
// @Success 			200 		{object}		domain.User
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/users/{id}	[GET]
func (c *UserController) ReadUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.svc.ReadUser(id)

	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateUser godoc
//
// @Description 	Update info user
// @Tags 					User
// @Accept 				json
// @Produce 			json
// @Param					id			path				string										true		"User ID"
// @Param					auth		body				UpdateUserRequest					true		"Auth admin"
// @Success 			200 		{object}		object{message=string}
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/users/{id}	[PATCH]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var json UpdateUserRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	if json.Email != nil {
		_, err := mail.ParseAddress(*json.Email)
		if err != nil {
			HandleError(ctx, http.StatusBadRequest, err)
			return
		}
	}

	user, err := c.svc.ReadUser(id)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.svc.UpdateUser(user.ID, ports.UpdateUserParams{
		Password: json.Password,
		FullName: json.FullName,
		Email:    json.Email,
	})
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

// DeleteUser godoc
//
// @Description 	Delete user
// @Tags 					User
// @Accept 				json
// @Produce 			json
// @Param					id			path				string								true		"User ID"
// @Success 			200 		{object}		object{message=string}
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/users/{id}	[DELETE]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.svc.ReadUser(id)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.svc.DeleteUser(user.ID)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
