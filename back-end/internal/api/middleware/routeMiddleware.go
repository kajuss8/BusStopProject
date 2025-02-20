package middleware

import (
	"github.com/gin-gonic/gin"
	"busProject/models"
	"net/http"
	"busProject/internal/database"
	)

func GetAllRoutes(ctx *gin.Context) {
	var routes []models.Route
	result := database.DB.Find(&routes)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"routesData": routes})
}

func GetRoutesWithDays(ctx *gin.Context) {
	routes, err := database.GetRoutesWithDays(database.DB)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Routes not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"routesData": routes})
}

