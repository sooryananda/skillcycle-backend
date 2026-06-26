package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/models"
)

func CreateListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
		Category    string  `json:"category" binding:"required"`
		Condition   string  `json:"condition" binding:"required"`
		ImageURL    string  `json:"image_url"`
		Location    string  `json:"location"`
		MarketDate  string  `json:"market_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketDate, err := time.Parse("2006-01-02", input.MarketDate)
	if err != nil {
		marketDate = time.Now()
	}

	listing := models.Listing{
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
		Condition:   input.Condition,
		ImageURL:    input.ImageURL,
		Location:    input.Location,
		MarketDate:  marketDate,
		IsAvailable: true,
	}

	if err := config.DB.Create(&listing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create listing"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Listing created successfully",
		"listing": listing,
	})
}

func GetAllListings(c *gin.Context) {
	var listings []models.Listing

	query := config.DB.Preload("User").Where("is_available = ?", true)

	// Filter by category if provided
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Filter by location if provided
	if location := c.Query("location"); location != "" {
		query = query.Where("location ILIKE ?", "%"+location+"%")
	}

	query.Order("created_at desc").Find(&listings)

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"count":    len(listings),
	})
}

func GetMyListings(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var listings []models.Listing
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&listings)

	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"count":    len(listings),
	})
}

func GetListingByID(c *gin.Context) {
	id := c.Param("id")

	var listing models.Listing
	if err := config.DB.Preload("User").First(&listing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"listing": listing})
}

func UpdateListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")

	var listing models.Listing
	if err := config.DB.First(&listing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	if listing.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own listings"})
		return
	}

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
		Condition   string  `json:"condition"`
		ImageURL    string  `json:"image_url"`
		Location    string  `json:"location"`
		IsAvailable bool    `json:"is_available"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&listing).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"message": "Listing updated successfully",
		"listing": listing,
	})
}

func DeleteListing(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id, _ := strconv.Atoi(c.Param("id"))

	var listing models.Listing
	if err := config.DB.First(&listing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Listing not found"})
		return
	}

	if listing.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own listings"})
		return
	}

	config.DB.Delete(&listing)

	c.JSON(http.StatusOK, gin.H{"message": "Listing deleted successfully"})
}
