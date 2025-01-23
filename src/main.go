package main

import (
	"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func getStopSchedule(ctx *gin.Context) {
	id := ctx.Param("id")
	schedule, err := models.CreateStopsSchedule(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Stop schedule": schedule})
}



func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	router.GET("/StopSchedle/:id", getStopSchedule)
	router.Run()
}
