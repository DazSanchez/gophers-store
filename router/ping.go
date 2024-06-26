package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// addPingRoutes registers RESTful controllers under /ping endpoint.
func addPingRoutes(rg *gin.RouterGroup) {
	r := rg.Group("ping")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
