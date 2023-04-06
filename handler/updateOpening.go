package handler

import (
	"fmt"
	"net/http"

	"github.com/LuisMarchio03/golang-boilerplate-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamsIsRequired("id",
			"queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf(
			"opening with id: %s not found",
			id,
		))
		return
	}

	// update opening
	if request.Role == "" {
		opening.Role = request.Role
	}
	if request.Company == "" {
		opening.Company = request.Company
	}
	if request.Location == "" {
		opening.Location = request.Location
	}
	if request.Link == "" {
		opening.Link = request.Link
	}
	if request.Remote == nil {
		opening.Remote = *request.Remote
	}
	if request.Salary <= 0 {
		opening.Salary = request.Salary
	}

	// save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)

}
