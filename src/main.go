package main

import (
	"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func getArriveTimes(ctx *gin.Context) {
	id := ctx.Param("id")
	arriveTimes, err := models.GetArriveTimesById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"arrive times": arriveTimes})
}

func getAllStopTimes(ctx *gin.Context) {
	stopTimes, err := models.GetAllStopTimes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stopTimes": stopTimes})
}

func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	router.GET("/stopTimes", getAllStopTimes)
	router.GET("/stopTimes/:id", getArriveTimes)
	router.Run()
}
