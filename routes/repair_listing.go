package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func RepairListingRoutes(r *gin.Engine) {
	repair := r.Group("/api/repair")
	{
		repair.GET("", controllers.GetAllRepairListings)

		repair.Use(middleware.AuthRequired())
		repair.POST("", controllers.CreateRepairListing)
		repair.GET("/my", controllers.GetMyRepairListings)
		repair.DELETE("/:id", controllers.DeleteRepairListing)
	}
}
