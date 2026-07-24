package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func InterestRoutes(r *gin.Engine) {
	interests := r.Group("/api/interests")
	interests.Use(middleware.AuthRequired())
	{
		interests.POST("", controllers.ToggleInterest)
		interests.GET("/my", controllers.GetMyInterests)
	}
}
