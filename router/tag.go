package router

import (
	"com.github.dazsanchez/gophers-store/controller/tag"
	"github.com/gin-gonic/gin"
)

func addTagRoutes(rg *gin.RouterGroup) {
	r := rg.Group("tags")

	r.GET("/", tag.FindAllController)
}
