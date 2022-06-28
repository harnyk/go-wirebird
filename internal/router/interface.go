package router

import "github.com/gin-gonic/gin"

type Router interface {
	AddLegacyEvent(c *gin.Context)
	AddEvent(c *gin.Context)
	HandleEventSocket(c *gin.Context)
}
