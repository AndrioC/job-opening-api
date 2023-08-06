package handler

import (
	"net/http"

	"github.com/andrioc/job-opening-api/schemas"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		logger.ErrorF("error validating data: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "creating-opening", opening)

}
