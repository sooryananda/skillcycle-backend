package models

import "time"

type RepairListing struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
	Title         string    `json:"title" gorm:"not null"`
	Description   string    `json:"description"`
	Skills        string    `json:"skills" gorm:"not null"`
	PriceRangeMin float64   `json:"price_range_min"`
	PriceRangeMax float64   `json:"price_range_max"`
	Location      string    `json:"location"`
	MarketDate    time.Time `json:"market_date"`
	IsAvailable   bool      `json:"is_available" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	SlotNumber    string    `json:"slot_number" gorm:"default:''"`
}
