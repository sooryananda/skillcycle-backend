package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/models"
)

func CreateRepairListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		Title         string   `json:"title" binding:"required"`
		Description   string   `json:"description"`
		Skills        []string `json:"skills" binding:"required"`
		PriceRangeMin float64  `json:"price_range_min"`
		PriceRangeMax float64  `json:"price_range_max"`
		Location      string   `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketDate := config.GetNextSunday()
	slotNumber := config.GenerateSlotNumber("repair")

	listing := models.RepairListing{
		UserID:        userID,
		Title:         input.Title,
		Description:   input.Description,
		Skills:        strings.Join(input.Skills, ", "),
		PriceRangeMin: input.PriceRangeMin,
		PriceRangeMax: input.PriceRangeMax,
		Location:      "Koramangala, Bangalore",
		MarketDate:    marketDate,
		SlotNumber:    slotNumber,
		IsAvailable:   true,
	}

	if err := config.DB.Create(&listing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create repair listing"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Repair listing created successfully",
		"listing": listing,
	})
}

func GetAllRepairListings(c *gin.Context) {
	var listings []models.RepairListing

	query := config.DB.Preload("User").Where("is_available = ?", true)

	if skill := c.Query("skill"); skill != "" {
		query = query.Where("skills ILIKE ?", "%"+skill+"%")
	}

	query.Order("created_at desc").Find(&listings)

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"count":    len(listings),
	})
}

func GetMyRepairListings(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var listings []models.RepairListing
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&listings)
	c.JSON(http.StatusOK, gin.H{"listings": listings, "count": len(listings)})
}

func DeleteRepairListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var listing models.RepairListing
	if err := config.DB.First(&listing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	if listing.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own listings"})
		return
	}

	config.DB.Delete(&listing)
	c.JSON(http.StatusOK, gin.H{"message": "Repair listing deleted"})
}
