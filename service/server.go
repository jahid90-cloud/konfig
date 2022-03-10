package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jahid90-cloud/konfig/service/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.GetPing)

	router.Run("localhost:8080")
}
