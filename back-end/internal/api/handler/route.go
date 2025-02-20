package handler

import (
	"github.com/gin-gonic/gin"
	"busProject/internal/api/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/bus-stops/v1/allStops", middleware.GetAllStops)
	r.GET("/api/bus-stops/v1/allCalendars", middleware.GetAllCalendars)
	r.GET("/api/bus-stops/v1/calendarById/:id", middleware.GetCalendarByID)
	r.GET("/api/bus-stops/v1/allRoutes", middleware.GetAllRoutes)
	r.GET("/api/bus-stops/v1/routesWithDays", middleware.GetRoutesWithDays)
	r.GET("/api/bus-stops/v1/allStopTimes", middleware.GetAllStopTimes)
	r.GET("/api/bus-stops/v1/allTrips", middleware.GetAllTrip)
}