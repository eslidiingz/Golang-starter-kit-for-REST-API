package migrations

import (
	"api-service-sdk/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Livestream{})
	// Add other model migrations here...
}
