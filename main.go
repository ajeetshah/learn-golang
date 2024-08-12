package main

import (
	"fmt"
	"log"
	"os"

	"example.com/learn-golang/database"
	"example.com/learn-golang/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var PORT string

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
	routes.SetupRoutes(r)
	r.Run(":" + PORT)
}
