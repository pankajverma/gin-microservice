package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-registry/constants"
)

func DiscoverHandler(context *gin.Context) {
	name := context.Param(constants.NameKey)
	for _, service := range services {
		if service.Name == name {
			context.JSON(http.StatusOK, gin.H{"host": service.Host, "port": service.Port})
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
}
