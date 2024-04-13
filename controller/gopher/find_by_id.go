package gopher

import (
	"net/http"
	"strconv"

	"com.github.dazsanchez/gophers-store/repository"
	"github.com/gin-gonic/gin"
)

func FindByIdController(ctx *gin.Context) {
	idStr := ctx.Params.ByName("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}

	g, err := repository.Gopher.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch gopher",
			"error":   err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, g)
}
