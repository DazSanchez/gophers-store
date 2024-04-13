package router

import (
	"com.github.dazsanchez/gophers-store/controller/gopher"
	"github.com/gin-gonic/gin"
)

func addGopherRoutes(rg *gin.RouterGroup) {
	r := rg.Group("gopher")

	r.GET("/:id", gopher.FindByIdController)
}
