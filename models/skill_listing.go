package models

import "time"

type SkillListing struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
	Title         string    `json:"title" gorm:"not null"`
	Description   string    `json:"description"`
	SkillType     string    `json:"skill_type" gorm:"not null"`
	Price         float64   `json:"price" gorm:"not null"`
	IsCustomOrder bool      `json:"is_custom_order" gorm:"default:false"`
	ImageURL      string    `json:"image_url"`
	Location      string    `json:"location"`
	MarketDate    time.Time `json:"market_date"`
	IsAvailable   bool      `json:"is_available" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	SlotNumber    string    `json:"slot_number" gorm:"default:''"`
}
