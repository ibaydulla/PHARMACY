package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
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

func MedicinesCreate(c *gin.Context) {

	var medicines models.Medicines

	if err := c.BindJSON(&medicines); err != nil {
		c.JSON(400, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	_, err := repositories.MedicinesCreate(c.Request.Context(), medicines)

	if err != nil {
		c.JSON(400, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
	})
}

func MedicinesDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.MedicinesDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

func MedicinesUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	var req models.Medicines

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.MedicinesUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, models.MedicinesErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}
	c.JSON(200, "ok")
}

func MedicinesRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/medicines", Medicineslist)
	rg.POST("/admin/medicines", MedicinesCreate)
	rg.DELETE("/admin/medicines/:id", MedicinesDelete)
	rg.PUT("/admin/medicines/:id", MedicinesUpdate)
}
