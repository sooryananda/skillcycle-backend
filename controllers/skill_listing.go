package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/models"
)

func CreateSkillListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		Title         string  `json:"title" binding:"required"`
		Description   string  `json:"description"`
		SkillType     string  `json:"skill_type" binding:"required"`
		Price         float64 `json:"price" binding:"required"`
		IsCustomOrder bool    `json:"is_custom_order"`
		ImageURL      string  `json:"image_url"`
		Location      string  `json:"location"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketDate := config.GetNextSunday()
	slotNumber := config.GenerateSlotNumber("skill")

	listing := models.SkillListing{
		UserID:        userID,
		Title:         input.Title,
		Description:   input.Description,
		SkillType:     input.SkillType,
		Price:         input.Price,
		IsCustomOrder: input.IsCustomOrder,
		ImageURL:      input.ImageURL,
		Location:      "Koramangala, Bangalore",
		MarketDate:    marketDate,
		SlotNumber:    slotNumber,
		IsAvailable:   true,
	}

	if err := config.DB.Create(&listing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create skill listing"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Skill listing created successfully",
		"listing": listing,
	})
}

func GetAllSkillListings(c *gin.Context) {
	var listings []models.SkillListing

	query := config.DB.Preload("User").Where("is_available = ?", true)

	if skillType := c.Query("skill_type"); skillType != "" {
		query = query.Where("skill_type = ?", skillType)
	}
	if customOnly := c.Query("custom_only"); customOnly == "true" {
		query = query.Where("is_custom_order = ?", true)
	}

	query.Order("created_at desc").Find(&listings)

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"count":    len(listings),
	})
}

func GetMySkillListings(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var listings []models.SkillListing
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&listings)
	c.JSON(http.StatusOK, gin.H{"listings": listings, "count": len(listings)})
}

func DeleteSkillListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var listing models.SkillListing
	if err := config.DB.First(&listing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	if listing.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own listings"})
		return
	}

	config.DB.Delete(&listing)
	c.JSON(http.StatusOK, gin.H{"message": "Skill listing deleted"})
}
