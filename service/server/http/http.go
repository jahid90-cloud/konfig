package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/server/http/config"
	"github.com/jahid90-cloud/konfig/service/server/http/middlewares"
)

type httpServer struct {
	router *gin.Engine
	config *config.Config
}

func NewServer() *httpServer {
	server := &httpServer{
		router: gin.New(),
		config: config.NewConfig(),
	}

	server.mountMiddlewares()
	server.mountRoutes()

	return server
}

func (s *httpServer) Start(addr string) error {
	if err := s.router.Run(addr); err != nil {
		return err
	}

	return nil
}

// mountMiddlewares Mounts global middlewares
func (s *httpServer) mountMiddlewares() {
	s.router.Use(gin.Logger(), gin.Recovery())
	s.router.Use(middlewares.PrimeRequestMiddleware())
}

// mountRoutes Mounts routers from each app
func (s *httpServer) mountRoutes() {
	for _, app := range s.config.Apps {
		app.Router(s.router)
	}
}