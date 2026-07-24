package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sooryananda/skillcycle-backend/config"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "SkillCycle backend is running!"})
	})

	// Register auth routes
	routes.AuthRoutes(r)

	routes.ListingRoutes(r)
	routes.SkillListingRoutes(r)
	routes.RepairListingRoutes(r)
	routes.MarketSlotRoutes(r)

	routes.InterestRoutes(r)

	// Demand pulse — proxies to Python AI service
	r.GET("/api/demand", controllers.GetDemandPulse)

	r.POST("/api/assess", controllers.GetAssessment)

	port := os.Getenv("PORT")
	log.Println("Server running on port", port)
	r.Run(":" + port)
}
