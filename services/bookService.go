package services

import (
	"errors"
	"net/http"

	"example.com/learn-golang/database"
	"example.com/learn-golang/models"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book *models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func ReadBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func ReadBooks(c *gin.Context) {
	var books []models.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	err := c.ShouldBind(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateBook models.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
}
