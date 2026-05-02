package handlers

import (
	"MarketPlace/mainProj/clients"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var reviewClient = clients.NewReviewClient()

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

func AddProductReview(c *gin.Context) {
	productIDStr := c.Param("id") // Это строка "3"

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

	input["product_id"] = productID

	result, err := reviewClient.AddReview(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
		return
	}
	c.JSON(http.StatusCreated, result)
}

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
