package middleware

import (
	"api-service-sdk/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-KEY")
		secretKey := c.GetHeader("SECRET-KEY")

		// Check if API key and Secret key are present
		if apiKey == "" || secretKey == "" {
			c.JSON(401, gin.H{
				"error":   true,
				"message": "API key and Secret key are invalid",
			})
			c.Abort()
			return
		}

		// Search for the user with the matching API key and Secret key
		var user models.User
		if err := db.Where("api_key = ? AND secret_key = ?", apiKey, secretKey).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Invalid API key or Secret key"})
			c.Abort()
			return
		}

		// Store the user object in the request context
		c.Set("user", user)

		c.Next()
	}
}
