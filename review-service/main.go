package main

import (
	"MarketPlace/review-service/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=reviews_db port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = database
	handlers.SetDB(DB)
	fmt.Println("ReviewService: Database connected successfully")
}

func main() {
	connectDatabase()

	r := gin.Default()

	r.GET("/reviews/:product_id", handlers.GetReviews)
	r.POST("/reviews", handlers.AddReview)
	r.DELETE("/reviews/:id", handlers.DeleteReview)

	r.Run(":8081")
}
