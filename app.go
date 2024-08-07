package main

import (
	"fmt"
	"log"
	"os"

	"example.com/learn-golang/controllers"
	"example.com/learn-golang/database"
	"example.com/learn-golang/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	PORT string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Coudn't load env file!!")
	}

	PORT = os.Getenv("PORT")
}

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()
	r := gin.Default()

	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/sign-up", controllers.CreateUser)
		publicRoutes.POST("/sign-in", controllers.BasicAuth, controllers.GetToken)
	}

	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middlewares.ValidateToken())
	{
		protectedRoutes.GET("/books/:id", controllers.ReadBook)
		protectedRoutes.GET("/books", controllers.ReadBooks)
		protectedRoutes.POST("/books", controllers.CreateBook)
		protectedRoutes.PUT("/books/:id", controllers.UpdateBook)
		protectedRoutes.DELETE("/books/:id", controllers.DeleteBook)
	}

	r.Run(":" + PORT)
}
