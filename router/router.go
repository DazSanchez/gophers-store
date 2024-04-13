package router

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	r := e.Group("api")

	addCategoryRoutes(r)
	addTagRoutes(r)
	addPingRoutes(r)
}
