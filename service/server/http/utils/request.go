package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/utils/logger"
)

func GetLoggerFromRequest(ctx *gin.Context) logger.Logger {
	l, available := ctx.Get("logger")
	if available {
		logger, ok := l.(logger.Logger)
		if ok {
			return logger
		}
	}

	return logger.NewNullLogger()
}
