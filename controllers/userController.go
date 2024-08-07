package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	services.CreateUser(c)
}
