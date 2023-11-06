package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinisioux/go-api/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First((&opening), id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
