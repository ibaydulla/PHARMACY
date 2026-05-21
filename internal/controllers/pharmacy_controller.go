package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
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

func PharmacyCreate(c *gin.Context) {

	var pharmacy models.Pharmacy

	if err := c.BindJSON(&pharmacy); err != nil {
		c.JSON(400, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	_, err := repositories.PharmacyCreate(c.Request.Context(), pharmacy)

	if err != nil {
		c.JSON(400, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
	})
}

func PharmacyDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.PharmacyDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

func PharmacyUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	var req models.Pharmacy

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.PharmacyUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, models.PharmacyErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}
	c.JSON(200, "ok")
}

func PharmacyRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/pharmacy", Pharmacylist)
	rg.POST("/admin/pharmacy", PharmacyCreate)
	rg.DELETE("/admin/pharmacy/:id", PharmacyDelete)
	rg.PUT("/admin/p/:id", PharmacyUpdate)
}
