package database

import (
	"fmt"
	"log"

	"example.com/learn-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "learn-golang"
	dbUser := "postgres"
	password := "password"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
