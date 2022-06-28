package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harnyk/go-wirebird/internal/router"
	"github.com/harnyk/go-wirebird/internal/wirebird"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	ginEngine := gin.Default()
	melodyRouter := melody.New()
	wb := wirebird.New(melodyRouter)
	appRoutes := router.New(melodyRouter, wb)

	ginEngine.POST("/api/updates", appRoutes.AddLegacyEvent)
	ginEngine.POST("/api/v2/events", appRoutes.AddEvent)
	ginEngine.GET("/api/v2/events-sock", appRoutes.HandleEventSocket)

	ginEngine.Run(":4380")
}
