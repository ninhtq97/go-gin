package restful

import (
	"fmt"
	"ninhtq/go-gin/core/exception"
	"ninhtq/go-gin/internal/ports"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeToken = "bearer"
	authorizationArgKey    = "authorization_arg"
	authorizationMeKey     = "authorization_me"
)

func (server *Server) parseToken(c *gin.Context) (ports.AuthDecoded, error) {
	authorizationHeader := c.GetHeader(authorizationHeaderKey)
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
	if authorizationType != authorizationTypeToken {
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
