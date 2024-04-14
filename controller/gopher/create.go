package gopher

import (
	"log"
	"net/http"

	"com.github.dazsanchez/gophers-store/dto"
	"com.github.dazsanchez/gophers-store/service"
	"github.com/gin-gonic/gin"
)

func CreateController(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't create gopher",
				"error":   r,
			})
		}
	}()

	var g dto.CreateGopherDTO
	err := ctx.BindJSON(&g)

	if err != nil {
		log.Panicln("invalid body: ", err)
	}

	response := service.Gopher.Create(g)

	ctx.JSON(http.StatusOK, response)
}
