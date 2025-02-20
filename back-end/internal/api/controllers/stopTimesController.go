package controllers

import (
	"github.com/gin-gonic/gin"
	"busProject/models"
	"net/http"
	"busProject/internal/database"
	)

func GetAllStopTimes(ctx *gin.Context) {
	var stopTimes []models.StopTime
	result := database.DB.Find(&stopTimes)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"stopTimesData": stopTimes})
}