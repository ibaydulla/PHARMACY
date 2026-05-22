package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
	"github.com/ibaydulla/internal/utils"
)

func Medicineslist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.Medicineslist(c, repositories.Medicinesfilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
})

	if err != nil {
	utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)
}

func MedicinesCreate(c *gin.Context) {

	var medicines models.Medicines

	if err := c.BindJSON(&medicines); 
	err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.MedicinesCreate(c.Request.Context(), medicines)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func MedicinesDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.MedicinesDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func MedicinesUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
	utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Medicines

	if err := c.BindJSON(&req); err != nil {
	utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.MedicinesUpdate(c.Request.Context(), id, req)
	if err != nil {
	utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func MedicinesRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/medicines", Medicineslist)
	rg.POST("/admin/medicines", MedicinesCreate)
	rg.DELETE("/admin/medicines/:id", MedicinesDelete)
	rg.PUT("/admin/medicines/:id", MedicinesUpdate)
}
