package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	services.Signin(c)
}
