package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for the presence of 'Secret' & 'API-Key' headers
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := c.GetHeader("APP-SECRET-KEY")
		apiKey := c.GetHeader("APP-API-KEY")

		secretKeyValid := os.Getenv("APP_SECRET_KEY")
		apiKeyValid := os.Getenv("APP_API_KEY")

		if secret != secretKeyValid || apiKey != apiKeyValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized, Access denied.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
