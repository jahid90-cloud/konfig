package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/server/http/config"
	"github.com/jahid90-cloud/konfig/service/server/http/middlewares"
	"github.com/jahid90-cloud/konfig/service/utils/env"
)

type httpServer struct {
	router *gin.Engine
	config *config.Config
}

func NewServer(env env.Env) *httpServer {

	engine := gin.New()
	config := config.NewConfig(env)

	// Remove when used behind a proxy
	engine.SetTrustedProxies(nil)

	server := &httpServer{
		router: engine,
		config: config,
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
	apiV1Router := s.router.Group("/api/v1")
	for _, app := range s.config.Apps {
		app.Router(apiV1Router)
	}
}
