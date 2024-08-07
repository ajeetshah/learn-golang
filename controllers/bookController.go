package controllers

import (
	"example.com/learn-golang/services"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	services.CreateBook(c)
}

func ReadBook(c *gin.Context) {
	services.ReadBook(c)
}

func ReadBooks(c *gin.Context) {
	services.ReadBooks(c)
}

func UpdateBook(c *gin.Context) {
	services.UpdateBook(c)
}

func DeleteBook(c *gin.Context) {
	services.DeleteBook(c)
}
