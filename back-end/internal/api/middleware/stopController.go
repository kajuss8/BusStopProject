package middleware

import (
	"github.com/gin-gonic/gin"
	"busProject/models"
	"net/http"
	"busProject/internal/database"
	)

func GetAllStops(ctx *gin.Context) {
	var stops []models.Stop
	result := database.DB.Find(&stops)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"stopsData": stops})
}