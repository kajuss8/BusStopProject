package middleware

import (
	"github.com/gin-gonic/gin"
	"busProject/models"
	"net/http"
	"busProject/internal/database"
	)

func GetAllTrip(ctx *gin.Context) {
	var trips []models.Trip
	result := database.DB.Find(&trips)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tripsData": trips})
}