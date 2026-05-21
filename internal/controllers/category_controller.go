package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
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
		c.JSON(500, gin.H{
			"success":   false,
			"error_msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
		"data":   list,
	})
}

func CategoryCreate(c *gin.Context) {

	var category models.Category

	if err := c.BindJSON(&category); err != nil {
		c.JSON(400, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	_, err := repositories.CategoryCreate(c.Request.Context(), category)

	if err != nil {
		c.JSON(400, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
	})
}

func CategoryDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.CategoryDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

func CategoryUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	var req models.Category

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.CategoryUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, models.CategoryErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}
	c.JSON(200, "ok")
}

func CategoryRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/category", Categorylist)
	rg.POST("/admin/category", CategoryCreate)
	rg.DELETE("/admin/category/:id", CategoryDelete)
	rg.PUT("/admin/category/:id", CategoryUpdate)
}
