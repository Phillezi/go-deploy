package middleware

import (
	"github.com/gin-gonic/gin"
	"go-deploy/pkg/sys"
	v1 "go-deploy/routers/api/v1"
	"go-deploy/service/user_service"
)

func SynchronizeUser(c *gin.Context) {
	context := sys.NewContext(c)

	auth, err := v1.WithAuth(&context)
	if err != nil {
		context.ServerError(err, v1.AuthInfoNotAvailableErr)
		return
	}

	_, err = user_service.New().WithAuth(auth).Create()
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	c.Next()
}
