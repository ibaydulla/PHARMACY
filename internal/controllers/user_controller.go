package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/repositories"
	"github.com/ibaydulla/internal/utils"
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
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, list)
}

func UserCreate(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); 
	err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	_, err := repositories.UserCreate(c.Request.Context(), user)

	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.UserDelete(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}

	utils.SuccessResponse(c, "")
}

func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		utils.ErrorResponse(c, err, 400, utils.ErrorCodeRequired)
		return
	}

	err = repositories.UserUpdate(c.Request.Context(), id, req)
	if err != nil {
	utils.ErrorResponse(c, err, 500, utils.ErrorCodeRequired)
		return
	}
	utils.SuccessResponse(c, "")
}

func UserRoute(rg *gin.RouterGroup) {
	rg.GET("/admin/users", Userlist)
	rg.POST("/admin/users", UserCreate)
	rg.DELETE("/admin/users/:id", UserDelete)
	rg.PUT("/admin/users/:id", UserUpdate)
}
