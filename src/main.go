package main

import (
	"busProject/src/handleFiles"
	"busProject/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "busProject/src/models"
// "net/http"

// "github.com/gin-gonic/gin"
// "busProject/src/handleFiles"

func getStopTimesByStopId(ctx *gin.Context) {
	id := ctx.Param("id")
	stopTimes, err := models.GetAllArriveTimesByStopId(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Stop not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stopTimes": stopTimes})
}

func main() {

	handleFiles.ProcessGtfs()

	router := gin.Default()
	router.GET("/stopTimes/:id", getStopTimesByStopId)
	router.Run()
}
