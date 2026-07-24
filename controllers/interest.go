package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/models"
)

func ToggleInterest(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		ListingID uint   `json:"listing_id" binding:"required"`
		Category  string `json:"category" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if interest already exists
	var existing models.Interest
	result := config.DB.Where(
		"user_id = ? AND listing_id = ?",
		userID, input.ListingID,
	).First(&existing)

	if result.Error == nil {
		// Already interested — remove it (toggle off)
		config.DB.Delete(&existing)
		c.JSON(http.StatusOK, gin.H{
			"message":    "Interest removed",
			"interested": false,
		})
		return
	}

	// Not yet interested — add it
	interest := models.Interest{
		UserID:    userID,
		ListingID: input.ListingID,
		Category:  input.Category,
	}
	config.DB.Create(&interest)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Interest registered",
		"interested": true,
	})
}

func GetMyInterests(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var interests []models.Interest
	config.DB.Where("user_id = ?", userID).Find(&interests)

	// Return just listing IDs for easy frontend lookup
	ids := make([]uint, len(interests))
	for i, interest := range interests {
		ids[i] = interest.ListingID
	}

	c.JSON(http.StatusOK, gin.H{"interested_listing_ids": ids})
}
