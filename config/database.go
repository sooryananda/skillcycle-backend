package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sooryananda/skillcycle-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto create tables
	database.AutoMigrate(
		&models.User{},
		&models.Listing{},
		&models.SkillListing{},
		&models.RepairListing{},
		&models.MarketSlot{},
		&models.Interest{},
	)

	log.Println("Database connected successfully!")
	DB = database
}
