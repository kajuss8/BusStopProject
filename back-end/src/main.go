package main

import (
	//"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func getStopSchedule(ctx *gin.Context) {
	id := ctx.Param("id")
	stopSchedule, err := models.CreateStopsSchedule(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stopSchedule": stopSchedule})
}

func getRouteSchedule(ctx *gin.Context) {
	id := ctx.Param("id")
	routeSchedule, err := models.CreateRouteSchedule(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"routeSchedules": routeSchedule})
}

func main() {

	//handleFiles.ProcessGtfs()

	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/StopSchedle/:id", getStopSchedule)
	router.GET("/RouteSchedule/:id", getRouteSchedule)
	router.Run()
}
