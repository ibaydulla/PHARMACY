package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
)

func Orderlist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	list, err := repositories.Userlist(c, repositories.Userfilter{
		Limit:  limit,
		Offset: offset,
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

func OrderCreate(c *gin.Context) {

	var user models.Order

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	_, err := repositories.OrderCreate(c.Request.Context(), user)

	if err != nil {
		c.JSON(400, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
	})
}

func OrderDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.OrderDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

func OrderUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	var req models.Order

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.OrderUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, models.OrderErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}
	c.JSON(200, "ok")
}

func OrderRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/orders", Orderlist)
	rg.POST("/admin/orders", OrderCreate)
	rg.DELETE("/admin/orders/:id", OrderDelete)
	rg.PUT("/admin/ordrs/:id", OrderUpdate)
}
