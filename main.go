package main

import (
	"api-service-sdk/database"
	"api-service-sdk/migrations"
	"api-service-sdk/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	println("[SUCCESS] LOAD ENV")

	// Init and Migrate Database
	db := database.InitDB()
	println("[SUCCESS] INIT DATABASE")
	migrations.AutoMigrate(db)
	println("[SUCCESS] MIGRATE DATABASE")

	r := gin.Default()
	routes.SetupAdminRoutes(r, db)
	routes.SetupRoutes(r, db)

	// Start the server
	r.Run(":8080")
}
