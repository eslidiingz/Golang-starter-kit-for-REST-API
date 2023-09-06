// controllers/user_controller.go
package controllers

import (
	"api-service-sdk/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	ProjectName string `json:"project_name" binding:"required"`
}

func CreateUser(db *gorm.DB, c *gin.Context) {
	var userReq UserRequest

	// Validate incoming request
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(400, gin.H{"error": "Missing or invalid fields"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password"})
		return
	}

	projectSlug := userReq.Username + "-" + strings.ReplaceAll(userReq.ProjectName, " ", "-")
	projectSlugLowerCase := strings.ToLower(projectSlug)

	// Populate the User model with the validated data
	var user models.User
	user.Username = userReq.Username
	user.Password = string(hashedPassword) // Store the hashed password
	user.Email = userReq.Email
	user.ProjectName = userReq.ProjectName
	user.ProjectSlug = projectSlugLowerCase
	user.SecretKey = strings.ReplaceAll(uuid.New().String(), "-", "")
	user.ApiKey = strings.ReplaceAll(uuid.New().String(), "-", "")

	// Save to DB and check for errors
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": true,
			"message": "Could not create user, This username {"+ userReq.Username + "} is already used",
		})
		return
	}

	// Successfully created
	c.JSON(200, gin.H{
		"success": true,
		"message": "User successfully created",
		"result":  user,
	})
}
