package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)

		// Protected
		auth.Use(middleware.AuthRequired())
		auth.GET("/profile", controllers.GetProfile)
		auth.PUT("/profile", controllers.UpdateProfile)
	}
}
