package router

import (
	"com.github.dazsanchez/gophers-store/controller/category"
	"github.com/gin-gonic/gin"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
	r := rg.Group("categories")

	r.GET("/", category.FindAllController)
}
