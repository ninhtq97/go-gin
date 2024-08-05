package controllers

import (
	"net/http"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/internal/ports"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc    ports.AuthService
	config config.Config
}

func NewAuthController(svc ports.AuthService, config config.Config) *AuthController {
	return &AuthController{
		svc:    svc,
		config: config,
	}
}

// Login godoc
//
// @Description 	Login
// @Tags 					Auth
// @Accept 				json
// @Produce 			json
// @Param					auth		body				LoginRequest					true		"Auth user"
// @Success 			200 		{object}		domain.LoginResponse
// @Failure				401			{object}		exception.Exception
// @Failure				422			{object}		exception.Exception
// @Router 				/login	[POST]
func (c *AuthController) Login(ctx *gin.Context) {
	var json LoginRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	authorized, err := c.svc.Login(ports.LoginArgs{
		Username: json.Username,
		Password: json.Password,
	})
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, authorized)
}
