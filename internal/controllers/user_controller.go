package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)
 //
func Userlist(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.Query("offset")
	offset, _ := strconv.Atoi(offsetStr)
	search := c.Query("search")
	role := c.Query("role")
	list, err := repositories.Userlist(c, repositories.Userfilter{
		Limit: limit,
		Offset: offset,
		Search: search,
		Role: role,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error_msg": err.Error(),
		})
		return
	} 

	c.JSON(200, gin.H{
		"succes": true,
		"data":   list,
	})
	return
}
func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("admin/users", Userlist)

}

//
func UserCreate(c *gin.Context) {

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.UserErrorResponse{err.Error(), "400"})
		return
	}

	_, err := repositories.UserCreate(c.Request.Context(), req)

	if err != nil {
		c.JSON(500, models.UserErrorResponse{err.Error(), "400"})
	}

	c.JSON(200, true)
}

//
func UserDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, err.Error())
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

// 
func UserUpdate(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	var req models.User

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = repositories.UserUpdate(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

// 
func UserRoutes(r *gin.Engine) {
	r.POST("/users", UserCreate)
	r.GET("/users", UserList)
	r.DELETE("/users/:id", UserDelete)
	r.PUT("/users/:id", UserUpdate)
}
