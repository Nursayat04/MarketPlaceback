package config

import (
	"fmt"
	"log"

	"MarketPlace/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	database.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{})

	DB = database
	fmt.Println("Database connected successfully")
}
