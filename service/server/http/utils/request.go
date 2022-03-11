package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/utils/logger"
)

var ErrMissingParamId = errors.New("missing param id")

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

func MapParamIdToConfiguration(ctx *gin.Context) (int64, error) {

	idStr := ctx.Param("id")

	if len(idStr) == 0 {
		return 0, ErrMissingParamId
	}

	id, err := strconv.ParseInt(idStr, 0, 0)
	if err != nil {
		return 0, err
	}

	return id, nil
}
