package services

import (
	"net/http"

	"example.com/learn-golang/database"
	"example.com/learn-golang/models"
	"example.com/learn-golang/utils"
	"github.com/gin-gonic/gin"
)

var tokens []string

func Signin(c *gin.Context) {
	var signInRequest models.SignInRequest
	if err := c.ShouldBindJSON(&signInRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var user models.User
	database.DB.Where("email=?", signInRequest.Email).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad credentials"})
		c.Abort()
		return
	}

	if !utils.IsPasswordMatch(user.Password, signInRequest.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad credentials"})
		c.Abort()
		return
	}

	token, _ := utils.GenerateJWT(user.Email, user.Role)
	tokens = append(tokens, token)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
