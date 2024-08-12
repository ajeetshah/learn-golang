package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func ReadUsers(c *gin.Context) {
	services.ReadUsers(c)
}
