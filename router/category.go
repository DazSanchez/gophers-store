package router

import (
	"com.github.dazsanchez/gophers-store/controller/category"
	"github.com/gin-gonic/gin"
)

// addCategoryRoutes registers RESTful controllers under /categories endpoint.
func addCategoryRoutes(rg *gin.RouterGroup) {
	r := rg.Group("categories")

	r.GET("/", category.FindAllController)
}
