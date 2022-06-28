package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router interface {
	AddLegacyEvent(c *gin.Context)
	AddEvent(c *gin.Context)
	HandleEventSocket(c *gin.Context)
	GetStaticFS() http.FileSystem
	GetIndexHTML() http.File
}
