package main

import (
	// "busProject/src/models"
	// "net/http"
	
	// "github.com/gin-gonic/gin"
	"busProject/src/handleFiles"
)

func main() {
	// router := gin.Default()
	// stops := models.CreateStops()
	// router.GET("/tasks", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"stops": stops})
	// })

	// router.Run()
	err := handleFiles.DownloadGtfs()
	if err != nil {
		panic(err)
	}

	err = handleFiles.Unzip()
	if err != nil {
		panic(err)
	}

}
