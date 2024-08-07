package database

import (
	"fmt"
	"log"
	"os"

	"example.com/learn-golang/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DB_HOST string
var DB_PORT string
var DB_NAME string
var DB_USER string
var DB_PASSWORD string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Coudn't load env file!!")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}

func DatabaseConnection() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_NAME,
		DB_PASSWORD,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Book{}, models.User{})
	if err != nil {
		log.Fatalln("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
