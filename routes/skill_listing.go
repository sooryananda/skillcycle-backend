package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func SkillListingRoutes(r *gin.Engine) {
	skills := r.Group("/api/skills")
	{
		skills.GET("", controllers.GetAllSkillListings)

		skills.Use(middleware.AuthRequired())
		skills.POST("", controllers.CreateSkillListing)
		skills.GET("/my", controllers.GetMySkillListings)
		skills.DELETE("/:id", controllers.DeleteSkillListing)
	}
}
