package services

import (
	"errors"
	"net/http"

	"example.com/learn-golang/database"
	"example.com/learn-golang/models"
	"example.com/learn-golang/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.ShouldBind(&user)
	user.Password = utils.GetPasswordHash(user.Password)
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

func ReadUsers(c *gin.Context) {
	var users []models.User
	res := database.DB.Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("users not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
