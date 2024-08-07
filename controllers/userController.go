package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	services.CreateUser(c)
}

func GetUser(c *gin.Context) {
	services.GetUser(c)
}
