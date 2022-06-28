package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/harnyk/go-wirebird/internal/models/compat"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	//Gin server with the following routes:
	//
	// - POST /api/updates - accepts a compat.LoggerEvent object and returns a 201 status code.
	//						Internally converts the event to models.LoggerEvent using models.Upgrade
	//						and prints it to stderr.
	// - POST /api/v2/events - accepts a models.LoggerEvent and returns 201 if successful. Prints the event to stderr
	r := gin.Default()
	m := melody.New()

	r.POST("/api/updates", func(c *gin.Context) {
		event := &compat.SerializedLoggerEvent{}
		err := c.BindJSON(event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		loggerEvent, err := models.Upgrade(event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("%+v\n", loggerEvent)

		c.JSON(http.StatusCreated, nil)

		log.Println(color.RedString("Compatibility mode. Some data may not be gathered. Please upgrade your client to the version >=2"))

		broadcastEvent, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshalling event: %s\n", err.Error())
			return
		}
		m.Broadcast(broadcastEvent)
	})

	r.POST("/api/v2/events", func(c *gin.Context) {
		var event models.LoggerEvent
		err := c.BindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("%+v\n", event)

		c.JSON(http.StatusCreated, nil)

		broadcastEvent, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshalling event: %s\n", err.Error())
			return
		}
		m.Broadcast(broadcastEvent)
	})

	r.GET("/api/v2/events-sock", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	r.Run(":4380")
}
