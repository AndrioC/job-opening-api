package handler

import (
	"net/http"

	"github.com/andrioc/job-opening-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		logger.ErrorF("error validating data: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	result := db.Model(&opening).Updates(request)

	if result.Error != nil {
		sendError(ctx, http.StatusBadRequest, "error when updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)

}
