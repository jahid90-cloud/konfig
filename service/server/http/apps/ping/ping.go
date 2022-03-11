package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/server/http/apps"
	"github.com/jahid90-cloud/konfig/service/server/http/utils"
)

type pingApp struct{}

func App() apps.App {
	return &pingApp{}
}

func (a *pingApp) Router(server *gin.Engine) *gin.RouterGroup {
	pingRouter := server.Group("/ping")
	{
		pingRouter.GET("/", a.pingHandler)
	}

	return pingRouter
}

// pingHandler Handles ping requests
func (a *pingApp) pingHandler(ctx *gin.Context) {
	utils.GetLoggerFromRequest(ctx).Info("handling ping request")
	ctx.String(http.StatusOK, "OK")
}
