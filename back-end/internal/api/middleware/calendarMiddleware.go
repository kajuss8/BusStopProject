package middleware

import (
	"busProject/internal/database"
	"busProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCalendars(ctx *gin.Context) {
	var calendar []models.Calendar
	result := database.DB.Find(&calendar)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"calendarsData": calendar})
}

func GetCalendarByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stop ID"})
		return
	}

	var calendar models.Calendar
	result := database.DB.First(&calendar, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Calendar not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"calendarData": calendar})
}