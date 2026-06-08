package controllers

import (
"net/http"
"github.com/gin-gonic/gin"

"github.com/ibaydulla/internal/models"
"github.com/ibaydulla/internal/repositories"
"github.com/ibaydulla/internal/utils"


)

func AuthRoute(rg *gin.RouterGroup) {

auth := rg.Group("/auth")

auth.POST("/register", Register)
auth.POST("/login", Login)
auth.POST("/logout", Logout)


}

func Register(c *gin.Context) {

var user models.User

if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	return
}

hashedPassword, err := utils.HashPassword(user.Password)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}

user.Password = hashedPassword

newUser, err := repositories.UserCreate(
	c.Request.Context(),
	user,
)

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}

c.JSON(http.StatusCreated, gin.H{
	"message": "user created",
	"user":    newUser,
})


}

func Login(c *gin.Context) {


var req models.LoginRequest

if err := c.ShouldBindJSON(&req); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	return
}

users, err := repositories.Userlist(
	c.Request.Context(),
	repositories.Userfilter{},
)

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}

var user models.User
found := false

for _, u := range users {

	if u.Email == req.Email {
		user = u
		found = true
		break
	}
}

if !found {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "user not found",
	})
	return
}

if err := utils.CheckPassword(
	req.Password,
	user.Password,
); err != nil {

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "wrong password",
	})
	return
}

token, err := utils.GenerateJWT(
	user.ID,
	user.Role,
)

if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
	return
}

c.JSON(http.StatusOK, gin.H{
	"message": "login successful",
	"token":   token,
})


}

func Logout(c *gin.Context) {


c.JSON(http.StatusOK, gin.H{
	"message": "logout successful",
})


}
