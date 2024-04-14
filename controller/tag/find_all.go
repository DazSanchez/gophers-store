package tag

import (
	"net/http"

	"com.github.dazsanchez/gophers-store/service"
	"github.com/gin-gonic/gin"
)

func FindAllController(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't fetch tags",
				"error":   r,
			})
		}
	}()

	tags := service.Tag.FindAll()

	ctx.JSON(http.StatusOK, tags)
}
