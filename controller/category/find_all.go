package category

import (
	"net/http"

	"com.github.dazsanchez/gophers-store/repository"
	"github.com/gin-gonic/gin"
)

func FindAllController(ctx *gin.Context) {
	categories, err := repository.Category.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch categories",
		})

		return
	}

	ctx.JSON(http.StatusOK, categories)
}
