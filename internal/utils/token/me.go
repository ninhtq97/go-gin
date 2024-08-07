package token

import (
	"ninhtq/go-gin/core/constants"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/core/exception"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) (*domain.User, error) {
	authorization, exists := c.Get(constants.AuthorizationMeKey)
	if !exists {
		return nil, exception.New(exception.TypePermissionDenied, "Unauthorized", nil)
	}
	me, ok := authorization.(*domain.User)
	if !ok {
		return nil, exception.New(exception.TypePermissionDenied, "Unauthorized", nil)
	}

	return me, nil
}
