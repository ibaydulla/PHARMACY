package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
)

func Userlist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	role := c.Query("role")
	list, err := repositories.Userlist(c, repositories.Userfilter{
		Limit:  limit,
		Offset: offset,
		Search: search,
		Role:   role,
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

func UserCreate(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	_, err := repositories.UserCreate(c.Request.Context(), user)

	if err != nil {
		c.JSON(400, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	c.JSON(200, gin.H{
		"succes": true,
	})
}

func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.UserDelete(c.Request.Context(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, "ok")
}

func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}

	err = repositories.UserUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, models.UserErrorResponse{
			Message: err.Error(),
			Code:    "400",
		})
		return
	}
	c.JSON(200, "ok")
}

func UserRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/users", Userlist)
	rg.POST("/admin/users", UserCreate)
	rg.DELETE("/admin/users/:id", UserDelete)
	rg.PUT("/admin/users/:id", UserUpdate)
}
