package apps

import "github.com/gin-gonic/gin"

// App An interface to be implemented by all http apps
type App interface {
	// Router Takes the server main router as input and attaches all its own sub-routes as a router group
	Router(server *gin.Engine) *gin.RouterGroup
}
