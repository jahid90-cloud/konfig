package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jahid90-cloud/konfig/service/utils/logger"
)

func PrimeRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := uuid.New().String()
		logger := logger.NewLogger()
		logger.SetLogLevel("INFO")
		logger.SetMetadata("(" + traceId + ")")

		c.Set("traceId", traceId)
		c.Set("logger", logger)
	}
}
