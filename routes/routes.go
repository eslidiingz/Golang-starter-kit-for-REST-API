package routes

import (
	"api-service-sdk/controllers"
	"api-service-sdk/middleware"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(db, c)
	})

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware(db))

	// Make a livestream route group
	liveStreamGroup := v1.Group("livestreams")
	{
		liveStreamGroup.GET("/", controllers.GetLiveStream)
		liveStreamGroup.POST("/", func(c *gin.Context) {
			controllers.CreateLivestream(db, c)
		})
	}

}
