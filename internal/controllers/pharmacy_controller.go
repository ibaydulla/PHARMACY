package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
	"github.com/ibaydulla/internal/utils"
)

func Pharmacylist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")

	list, err := repositories.Pharmacylist(c, repositories.Pharmacyfilter{
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

func PharmacyCreate(c *gin.Context) {

	var pharmacy models.Pharmacy

	if err := c.BindJSON(&pharmacy); 
	err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.PharmacyCreate(c.Request.Context(), pharmacy)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func PharmacyDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Pharmacy

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.PharmacyUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func PharmacyRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacy", Pharmacylist)
	rg.POST("/admin/pharmacy", PharmacyCreate)
	rg.DELETE("/admin/pharmacy/:id", PharmacyDelete)
	rg.PUT("/admin/p/:id", PharmacyUpdate)
}
