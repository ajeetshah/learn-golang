package main

import (
	"fmt"

	"example.com/learn-golang/controllers"
	"example.com/learn-golang/database"
	"example.com/learn-golang/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.POST("/login", controllers.BasicAuth, controllers.GetToken)

	r.Use(middlewares.ValidateToken())

	r.GET("/books/:id", controllers.ReadBook)
	r.GET("/books", controllers.ReadBooks)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run(":5000")
}
