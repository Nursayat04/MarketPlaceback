package handlers

import (
	"MarketPlace/mainProj/clients"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var reviewClient = clients.NewReviewClient()

// GET /products/:id/reviews
func GetProductReviews(c *gin.Context) {
	productID := c.Param("id")

	result, err := reviewClient.GetReviews(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	var data interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from ReviewService"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// POST /products/:id/reviews
func AddProductReview(c *gin.Context) {
	productIDStr := c.Param("id") // Это строка "3"

	// 1. Конвертируем строку в число
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 2. Кладем в map ЧИСЛО, а не строку
	input["product_id"] = productID

	// 3. Отправляем в Review Service
	result, err := reviewClient.AddReview(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
		return
	}

	// (Дополнительно) Проверь, чтобы ты не возвращал 201, если в result есть ошибка
	c.JSON(http.StatusCreated, result)
}

// DELETE /reviews/:id
func DeleteProductReview(c *gin.Context) {
	id := c.Param("id")

	result, err := reviewClient.DeleteReview(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	var data interface{}
	if err := json.Unmarshal([]byte(result), &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from ReviewService"})
		return
	}

	c.JSON(http.StatusOK, data)
}
