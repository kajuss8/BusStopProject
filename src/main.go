package main

import (
	"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTransportArrivalTimes(ctx *gin.Context) {
	id := ctx.Param("id")
	arriveTimes, err := models.GetArrivalTimes(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"arrival times": arriveTimes})
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

func getStopName(ctx *gin.Context) {
	id := ctx.Param("id")
	stopTimes, err := models.GetStopName(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stop name": stopTimes})
}

func getTripId(ctx *gin.Context) {
	id := ctx.Param("id")
	tripId, err := models.GetTripId(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tripId": tripId})
}

func getWeekWorkDays(ctx *gin.Context){
	id := ctx.Param("id")
	tempId, _ := strconv.Atoi(id)
	tripId, err := models.GetCalendarWorkDays(tempId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No such service ID"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Wrok days": tripId})
}

func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	router.GET("/tripId/:id", getTripId)
	router.GET("/stopsName/:id", getStopName)
	router.GET("/stopTimesById/:id", getStopTimesById)
	router.GET("/arrivalTimes/:id", getTransportArrivalTimes)
	router.GET("/stop/:id", getStopById)
	router.GET("/WorkDays/:id", getWeekWorkDays)
	router.Run()
}
