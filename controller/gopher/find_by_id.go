package gopher

import (
	"log"
	"net/http"
	"strconv"

	"com.github.dazsanchez/gophers-store/service"
	"github.com/gin-gonic/gin"
)

func FindByIdController(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't fetch gopher",
				"error":   r,
			})
		}
	}()

	idStr := ctx.Params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		log.Panicln("invalid id format")
	}

	g := service.Gopher.FindById(id)

	ctx.JSON(http.StatusOK, g)
}
