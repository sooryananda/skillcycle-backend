package config

import (
	"fmt"
	"time"
)

func GetNextSunday() time.Time {
	now := time.Now()
	daysUntilSunday := (7 - int(now.Weekday())) % 7
	if daysUntilSunday == 0 {
		daysUntilSunday = 7
	}
	next := now.AddDate(0, 0, daysUntilSunday)
	return time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, time.Local)
}

func GenerateSlotNumber(listingType string) string {
	var count int64

	switch listingType {
	case "skill":
		DB.Raw("SELECT COUNT(*) FROM skill_listings").Scan(&count)
		return fmt.Sprintf("SK%02d", count+1)
	case "repair":
		DB.Raw("SELECT COUNT(*) FROM repair_listings").Scan(&count)
		return fmt.Sprintf("RP%02d", count+1)
	default:
		DB.Raw("SELECT COUNT(*) FROM listings").Scan(&count)
		return fmt.Sprintf("SL%02d", count+1)
	}
}
