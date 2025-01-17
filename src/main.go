package main

import (
	"busProject/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "busProject/src/models"
// "net/http"

// "github.com/gin-gonic/gin"
// "busProject/src/handleFiles"

func main() {

	// err := handleFiles.DownloadGtfs()
	// if err != nil {
	// 	panic(err)
	// }

	// err = handleFiles.Unzip()
	// if err != nil {
	// 	panic(err)
	// }

	router := gin.Default()
	router.GET("/stopTimeByStopId/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		stopTimes, err := models.GetAllArriveTimesByStopId(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"stopTimes": stopTimes})	
	})

	router.Run()
}
