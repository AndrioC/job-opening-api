package handler

import (
	"net/http"

	"github.com/andrioc/job-opening-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listininge erros")
	}

	sendSuccess(ctx, "list-openings", openings)
}
