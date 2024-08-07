package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func BasicAuth(c *gin.Context) {
	services.BasicAuth(c)
}

func GetToken(c *gin.Context) {
	services.GetToken(c)
}
