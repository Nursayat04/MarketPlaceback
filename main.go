package main

import (
	"MarketPlace/config"
	"MarketPlace/handlers"
	"MarketPlace/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/products", handlers.GetProducts)
		protected.POST("/products", handlers.AddProduct)
		protected.GET("/products/:id", handlers.GetProductByID)
		protected.PUT("/products/:id", handlers.UpdateProduct)
		protected.DELETE("/products/:id", handlers.DeleteProduct)

		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.AddCategory)
		protected.DELETE("/categories/:id", handlers.DeleteCategory)

		protected.GET("/users", handlers.GetUsers)
		protected.DELETE("/users/:id", handlers.DeleteUser)

		protected.GET("/products/:id/reviews", handlers.GetProductReviews)
		protected.POST("/products/:id/reviews", handlers.AddProductReview)
		protected.DELETE("/reviews/:id", handlers.DeleteProductReview)
	}

	// Запуск сервера
	r.Run(":8080")
}
