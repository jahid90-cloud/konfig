package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}
