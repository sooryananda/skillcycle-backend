package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/models"
)

func generateSlotNumber(role string) string {
	var count int64
	config.DB.Model(&models.MarketSlot{}).Count(&count)
	prefix := "S"
	if role == "repair_person" {
		prefix = "R"
	}
	return fmt.Sprintf("%s%02d", prefix, count+1)
}

func BookMarketSlot(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var input struct {
		Role        string `json:"role" binding:"required"`
		MarketDate  string `json:"market_date"`
		Location    string `json:"location"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketDate, err := time.Parse("2006-01-02", input.MarketDate)
	if err != nil {
		marketDate = time.Now()
	}

	slotNumber := generateSlotNumber(input.Role)

	slot := models.MarketSlot{
		UserID:      userID,
		Role:        input.Role,
		MarketDate:  marketDate,
		Location:    input.Location,
		Description: input.Description,
		SlotNumber:  slotNumber,
		Status:      "confirmed",
	}

	if err := config.DB.Create(&slot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not book slot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Slot booked successfully",
		"slot":        slot,
		"slot_number": slotNumber,
	})
}

func GetAllMarketSlots(c *gin.Context) {
	var slots []models.MarketSlot
	config.DB.Preload("User").Where("status = ?", "confirmed").
		Order("created_at desc").Find(&slots)

	c.JSON(http.StatusOK, gin.H{
		"slots": slots,
		"count": len(slots),
	})
}

func GetMySlots(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var slots []models.MarketSlot
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&slots)
	c.JSON(http.StatusOK, gin.H{"slots": slots})
}
