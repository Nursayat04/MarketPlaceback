package main

import (
	"MarketPlace/mainProj/config"
	handlers2 "MarketPlace/mainProj/handlers"
	"MarketPlace/mainProj/middleware"
	"MarketPlace/mainProj/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{})
	if err != nil {
		fmt.Println("Migration failed:", err)
	}
	r := gin.Default()

	r.POST("/register", handlers2.Register)
	r.POST("/login", handlers2.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", handlers2.GetProducts)
		protected.POST("/products", handlers2.AddProduct)
		protected.GET("/products/:id", handlers2.GetProductByID)
		protected.PUT("/products/:id", handlers2.UpdateProduct)
		protected.DELETE("/products/:id", handlers2.DeleteProduct)

		protected.GET("/categories", handlers2.GetCategories)
		protected.POST("/categories", handlers2.AddCategory)
		protected.DELETE("/categories/:id", handlers2.DeleteCategory)

		protected.GET("/users", handlers2.GetUsers)
		protected.DELETE("/users/:id", handlers2.DeleteUser)

		protected.GET("/products/:id/reviews", handlers2.GetProductReviews)
		protected.POST("/products/:id/reviews", handlers2.AddProductReview)
		protected.DELETE("/reviews/:id", handlers2.DeleteProductReview)
	}

	r.Run(":8080")
}
