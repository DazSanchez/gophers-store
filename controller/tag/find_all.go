package tag

import (
	"net/http"

	"com.github.dazsanchez/gophers-store/repository"
	"github.com/gin-gonic/gin"
)

func FindAllController(ctx *gin.Context) {
	tags, err := repository.Tag.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch tags",
		})

		return
	}

	ctx.JSON(http.StatusOK, tags)
}
