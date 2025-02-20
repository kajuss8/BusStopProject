package routes

import (
	"github.com/gin-gonic/gin"
	"busProject/internal/api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/bus-stops/v1/allStops", controllers.GetAllStops)
	r.GET("/api/bus-stops/v1/allCalendars", controllers.GetAllCalendars)
	r.GET("/api/bus-stops/v1/calendarById/:id", controllers.GetCalendarByID)
	r.GET("/api/bus-stops/v1/allRoutes", controllers.GetAllRoutes)
	r.GET("/api/bus-stops/v1/routesWithDays", controllers.GetRoutesWithDays)
	r.GET("/api/bus-stops/v1/allStopTimes", controllers.GetAllStopTimes)
	r.GET("/api/bus-stops/v1/allTrips", controllers.GetAllTrip)
}