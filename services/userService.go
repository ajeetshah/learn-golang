package services

import (
	"net/http"

	"example.com/learn-golang/database"
	"example.com/learn-golang/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBind(&user)
	user.Password = GetPasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(user)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}

func GetPasswordHash(password string) string {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}
