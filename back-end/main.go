package main

import (
	"busProject/internal/config"
	"log"
	"github.com/gin-gonic/gin"
	"busProject/internal/database"
	"busProject/internal/api/handler"
)

const(
	configPath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/back-end/internal/config"
	configName = "config"
	ConfigType = "yaml"
)

// func corsMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

// func getStopSchedule(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stop ID"})
// 		return
// 	}

// 	stopSchedule, err := models.CreateStopsSchedule(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"stopSchedule": stopSchedule})
// }

// func getRouteSchedule(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	routeSchedule, err := models.CreateRouteSchedule(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"routeSchedules": routeSchedule})
// }

// func getAllRoutes(ctx *gin.Context) {
// 	routes, err := models.CreateRouteWorkDays()
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Routes not found"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"routesData": routes})
// }

// func getAllStops(ctx *gin.Context) {
// 	stops, err := models.GetAllStops()
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stops not found"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"stopsData": stops})
// }

// var db *gorm.DB

const(
	filepath = "C:/Users/Kajus.Sciaponis/Desktop/BusStopProject/back-end/internal/gtfs/gtfsFolder/"
	stopFileName = "stops.txt"
	routeFileName = "routes.txt"
	CalendarFileName = "calendar.txt"
	stopTimeFileName = "stop_times.txt"
	tripFileName = "trips.txt"
)

func main() {

	config, err := config.LoadConfiguration(configPath, configName, ConfigType)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	

	r := gin.Default()
	database.InitDb(config.DbPath)
	handler.RegisterRoutes(r)

	// 	router.Use(corsMiddleware())
	// 	router.GET("/AllRoutes", getAllRoutes)
	// 	router.GET("/StopSchedle/:id", getStopSchedule)
	// 	router.GET("/RouteSchedule/:id", getRouteSchedule)
	// router.GET("/stops", getAllStops)
	// router.GET("/calendars", getAllCalendars)
	// router.GET("/routes", getAllRoutes)
	// router.GET("/stopTimes", getAllStopTimes)
	// router.GET("/trips", getAllTrip)
	r.Run(":8080")
}
