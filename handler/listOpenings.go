package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinisioux/go-api/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error while fetching openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
