package main

import (
	"MarketPlace/review-service/handlers"
	"MarketPlace/review-service/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "host=db user=postgres password=postgres dbname=reviews_db port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	err = database.AutoMigrate(&models.Review{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	handlers.SetDB(DB)
	fmt.Println("ReviewService: Database connected and table created successfully")
}

func main() {
	connectDatabase()

	r := gin.Default()

	r.GET("/reviews/:product_id", handlers.GetReviews)
	r.POST("/reviews", handlers.AddReview)
	r.DELETE("/reviews/:id", handlers.DeleteReview)

	r.Run(":8081")
}
