package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
	"github.com/ibaydulla/internal/utils"
)

func Orderlist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	list, err := repositories.Orderlist(c, repositories.Orderfilter{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)
}

func OrderCreate(c *gin.Context) {

	var user models.Order

	if err := c.BindJSON(&user); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.OrderCreate(c.Request.Context(), user)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func OrderDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.OrderDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func OrderUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Order

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.OrderUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func OrderRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/orders", Orderlist)
	rg.POST("/admin/orders", OrderCreate)
	rg.DELETE("/admin/orders/:id", OrderDelete)
	rg.PUT("/admin/ordrs/:id", OrderUpdate)
}
