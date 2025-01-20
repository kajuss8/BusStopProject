package main

import (
	"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRoutes(ctx *gin.Context) {
	id := ctx.Param("id")
	trips, _ := models.GetAllTrips()
	stopTime, _ := models.GetStopTimesByStopId(id)
	tripIds := models.GetTripIds(stopTime)
	routeIds := models.GetTripsByIds(tripIds, trips)
	// if err != nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"arrival times": routeIds})
}

func getDifferentRouts(ctx *gin.Context) {
	id := ctx.Param("id")
	trips, _ := models.GetAllTrips()
	stopTime, _ := models.GetStopTimesByStopId(id)
	tripIds := models.GetTripIds(stopTime)
	routeIds := models.GetTripsByIds(tripIds, trips)
	routeonlyIds := models.GetRouteIds(routeIds)

	routs, _ := models.GetAllRoutes()

	diffRouts := models.GetDifferentRouts(routeonlyIds, routs)
	ctx.JSON(http.StatusOK, gin.H{"arrival times": diffRouts})

}

func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	//router.GET("/StopSchedle/:id", getStopSchedule)
	router.GET("/RouteIds/:id", getRoutes)
	router.GET("/differentRouts/:id", getDifferentRouts)
	router.Run()
}
