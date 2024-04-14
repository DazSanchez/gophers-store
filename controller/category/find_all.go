package category

import (
	"net/http"

	"com.github.dazsanchez/gophers-store/service"
	"github.com/gin-gonic/gin"
)

// FindAllController returns all available Categories in the database.
// It panics if there's a problem while retrieving data.
func FindAllController(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't fetch categories",
				"error":   r,
			})
		}
	}()

	cs := service.Category.FindAll()

	ctx.JSON(http.StatusOK, cs)
}
