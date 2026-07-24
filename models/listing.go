package models

import "time"

type Listing struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Category    string    `json:"category" gorm:"not null"`
	Condition   string    `json:"condition" gorm:"not null"`
	ImageURL    string    `json:"image_url"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"is_available" gorm:"default:true"`
	MarketDate  time.Time `json:"market_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SlotNumber  string    `json:"slot_number" gorm:"default:''"`
}
