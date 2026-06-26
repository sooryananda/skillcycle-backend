package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func ListingRoutes(r *gin.Engine) {
	listings := r.Group("/api/listings")
	{
		// Public — anyone can view listings
		listings.GET("", controllers.GetAllListings)
		listings.GET("/:id", controllers.GetListingByID)

		// Protected — must be logged in
		listings.Use(middleware.AuthRequired())
		listings.POST("", controllers.CreateListing)
		listings.GET("/my/listings", controllers.GetMyListings)
		listings.PUT("/:id", controllers.UpdateListing)
		listings.DELETE("/:id", controllers.DeleteListing)
	}
}
