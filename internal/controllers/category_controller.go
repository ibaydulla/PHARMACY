package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
	"github.com/ibaydulla/internal/utils"
)

func Categorylist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	list, err := repositories.Categorylist(c, repositories.Categoryfilter{
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

func CategoryCreate(c *gin.Context) {

	var category models.Category

	if err := c.BindJSON(&category); 
	err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.CategoryCreate(c.Request.Context(), category)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func CategoryDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.CategoryDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
	
}

func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.Category

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.CategoryUpdate(c.Request.Context(), id, req)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func CategoryRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/category", Categorylist)
	rg.POST("/admin/category", CategoryCreate)
	rg.DELETE("/admin/category/:id", CategoryDelete)
	rg.PUT("/admin/category/:id", CategoryUpdate)
}
