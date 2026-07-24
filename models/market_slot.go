package models

import "time"

type MarketSlot struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	Role        string    `json:"role" gorm:"not null"`
	MarketDate  time.Time `json:"market_date"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	SlotNumber  string    `json:"slot_number"`
	Status      string    `json:"status" gorm:"default:'confirmed'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
