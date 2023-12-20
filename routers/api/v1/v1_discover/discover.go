package v1_discover

import (
	"github.com/gin-gonic/gin"
	"go-deploy/pkg/sys"
	v1 "go-deploy/routers/api/v1"
	"go-deploy/service/discover_service"
)

func Discover(c *gin.Context) {
	context := sys.NewContext(c)

	discover, err := discover_service.Discover()
	if err != nil {
		context.ServerError(err, v1.InternalError)
		return
	}

	context.Ok(discover.ToDTO())
}