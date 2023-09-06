package routes

import (
	"api-service-sdk/controllers"
	"api-service-sdk/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminRoutes(r *gin.Engine, db *gorm.DB) {
	adminV1 := r.Group("/api-admin/v1/")
	adminV1.Use(middleware.AdminMiddleware())

	// Make a user route group
	userGroup := adminV1.Group("users")
	{
		userGroup.POST("/", func(c *gin.Context) {
			controllers.CreateUser(db, c)
		})
	}
}
