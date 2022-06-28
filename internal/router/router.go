package router

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/harnyk/go-wirebird/internal/models/compat"
	"github.com/harnyk/go-wirebird/internal/wirebird"
	"gopkg.in/olahol/melody.v1"
)

//go:embed dist/*
var webui embed.FS

type router struct {
	melody   *melody.Melody
	wirebird wirebird.Wirebird
}

func New(melody *melody.Melody, wirebird wirebird.Wirebird) Router {
	r := &router{
		melody:   melody,
		wirebird: wirebird,
	}

	return r
}

func (r *router) AddLegacyEvent(c *gin.Context) {
	event := &compat.SerializedLoggerEvent{}
	err := c.BindJSON(event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
	err = r.wirebird.BroadcastLegacyEvent(event)
	if err != nil {
		log.Println(color.RedString("Error broadcasting event: %s", err.Error()))
	}

	log.Println(color.RedString("Compatibility mode. Some data may not be gathered. Please upgrade your client to the version >=2"))
}

func (r *router) AddEvent(c *gin.Context) {
	event := &models.LoggerEvent{}
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)

	err = r.wirebird.BroadcastEvent(event)
	if err != nil {
		log.Println(color.RedString("Error broadcasting event: %s", err.Error()))
	}
}

func (r *router) HandleEventSocket(c *gin.Context) {
	r.melody.HandleRequest(c.Writer, c.Request)
}

func (r *router) GetStaticFS() http.FileSystem {
	sub, err := fs.Sub(webui, "dist")
	if err != nil {
		log.Fatalf("Error getting static files: %s", err.Error())
	}
	return http.FS(sub)
}
