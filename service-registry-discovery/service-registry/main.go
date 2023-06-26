package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"service-registry/handler"
)

func main() {
	router := gin.Default()

	router.Group("v1", func(context *gin.Context) {
		router.POST("/register", handler.RegisterHandler)
		router.GET("/discover/:name", handler.DiscoverHandler)

	})

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Failed to start service-registry:", err)
	}
}
