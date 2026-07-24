package models

import "time"

type Interest struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	ListingID uint      `json:"listing_id" gorm:"not null"`
	Category  string    `json:"category" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
