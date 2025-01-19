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

func getStopById(ctx *gin.Context) {
	id := ctx.Param("id")
	stop, err := models.GetStopById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stop": stop})
}

func getStopTimesById(ctx *gin.Context) {
	id := ctx.Param("id")
	stopTimes, err := models.GetStopTimesByStopId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stop Times": stopTimes})
}

func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	router.GET("/stopTimesById/:id", getStopTimesById)
	router.GET("/arriveTimes/:id", getArriveTimes)
	router.GET("/stop/:id", getStopById)
	router.Run()
}
