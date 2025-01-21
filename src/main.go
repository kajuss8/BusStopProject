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
	routeonlyIds := models.GetTripRouteIds(routeIds)

	routs, _ := models.GetAllRoutes()

	diffRouts := models.GetDifferentRouts(routeonlyIds, routs)
	ctx.JSON(http.StatusOK, gin.H{"arrival times": diffRouts})

}

func getRouteIds(ctx *gin.Context) {
	id := ctx.Param("id")
	trips, _ := models.GetAllTrips()
	stopTime, _ := models.GetStopTimesByStopId(id)
	tripIds := models.GetTripIds(stopTime)
	tripsById := models.GetTripsByIds(tripIds, trips)

	routeIds := models.GetMapTripsShapeRouteId(tripsById)
	ctx.JSON(http.StatusOK, gin.H{"Routes": routeIds})
}

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
	router.GET("/RouteIds/:id", getRoutes)
	router.GET("/differentRouts/:id", getDifferentRouts)
	router.GET("/routtes/:id", getRouteIds)
	router.Run()
}
