package restful

import (
	"errors"
	"fmt"
	"net/http"
	"ninhtq/go-gin/core/constants"
	"ninhtq/go-gin/core/domain"
	"ninhtq/go-gin/core/exception"
	"ninhtq/go-gin/internal/controllers"
	"ninhtq/go-gin/internal/ports"
	"ninhtq/go-gin/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func (server *Server) parseToken(c *gin.Context) (ports.AuthDecoded, error) {
	authorizationHeader := c.GetHeader(constants.AuthorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		msg := "authorization header not provided"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return ports.AuthDecoded{}, err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		msg := "invalid authorization format"
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return ports.AuthDecoded{}, err
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != constants.AuthorizationTypeToken {
		msg := fmt.Sprintf("authorization type %s not supported", authorizationType)
		err := exception.New(exception.TypePermissionDenied, msg, nil)
		return ports.AuthDecoded{}, err
	}

	token := fields[1]
	payload, err := server.service.TokenMaker().VerifyToken(token)
	if err != nil {
		return ports.AuthDecoded{}, err
	}
	return ports.AuthDecoded{Token: token, Payload: payload}, nil
}

func (server *Server) VerifyUserToken(token string) (*domain.User, error) {
	payload, err := server.service.TokenMaker().VerifyToken(token)
	if err != nil {
		return nil, err
	}

	user, err := server.service.User().ReadUser(payload.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (server *Server) VerifyUserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := controllers.AuthHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			utils.HandleError(c, http.StatusUnauthorized, err)
			return
		}

		if h.Token == nil {
			utils.HandleError(c, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		user, err := server.VerifyUserToken(*h.Token)
		if err != nil {
			utils.HandleError(c, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		c.Set(constants.AuthorizationMeKey, user)
		c.Next()
	}
}
