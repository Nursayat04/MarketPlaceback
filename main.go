package main

import (
	"MarketPlace/config"
	"MarketPlace/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.AddUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)

	r.GET("/products", handlers.GetProducts)
	r.POST("/products", handlers.AddProduct)
	r.GET("/products/:id", handlers.GetProductByID)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.Run(":8080")
}
