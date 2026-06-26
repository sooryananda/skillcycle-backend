package models

import "time"

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string    `json:"name" gorm:"not null"`
	Email         string    `json:"email" gorm:"unique;not null"`
	Password      string    `json:"-" gorm:"not null"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role" gorm:"default:'browser'"`
	SecondaryRole string    `json:"secondary_role" gorm:"default:''"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
