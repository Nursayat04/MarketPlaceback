package handlers

import (
	"MarketPlace/review-service/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetDB вызывается из mainProj.go для передачи подключения к БД
func SetDB(database *gorm.DB) {
	db = database
}

// GET /reviews/:product_id
func GetReviews(c *gin.Context) {
	productIDStr := c.Param("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	var reviews []models.Review
	if err := db.Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// POST /reviews
func AddReview(c *gin.Context) {
	var input struct {
		ProductID uint   `json:"product_id" binding:"required"`
		UserID    uint   `json:"user_id"    binding:"required"`
		Username  string `json:"username"   binding:"required"`
		Rating    int    `json:"rating"     binding:"required"`
		Comment   string `json:"comment"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Rating < 1 || input.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	review := models.Review{
		ProductID: input.ProductID,
		UserID:    input.UserID,
		Username:  input.Username,
		Rating:    input.Rating,
		Comment:   input.Comment,
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// DELETE /reviews/:id
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	result := db.Delete(&models.Review{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
