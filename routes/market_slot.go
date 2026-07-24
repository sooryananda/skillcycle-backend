package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sooryananda/skillcycle-backend/controllers"
	"github.com/sooryananda/skillcycle-backend/middleware"
)

func MarketSlotRoutes(r *gin.Engine) {
	slots := r.Group("/api/slots")
	{
		slots.GET("", controllers.GetAllMarketSlots)

		slots.Use(middleware.AuthRequired())
		slots.POST("", controllers.BookMarketSlot)
		slots.GET("/my", controllers.GetMySlots)
	}
}
