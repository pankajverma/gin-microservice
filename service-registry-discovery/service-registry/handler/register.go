package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"service-registry/models"
)

var services []models.Service

func RegisterHandler(context *gin.Context) {
	var service models.Service
	err := context.ShouldBindJSON(&service)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register service"})
		return
	}
	services = append(services, service)
	fmt.Printf("Registered service %s - %s: %d \n", service.Name, service.Host, service.Port)
	context.JSON(http.StatusOK, gin.H{"message": "Service registered successfully"})
}
