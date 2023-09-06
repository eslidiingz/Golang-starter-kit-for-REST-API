package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func InitDB() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	dbport := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("DATABASE_SSL_MODE")
	
	dsn := "host="+ host +" user="+ user +" dbname="+ dbname +" sslmode="+ sslmode +" password="+ password +" port="+ dbport +" TimeZone=Asia/Bangkok"
	// dsn := "host=localhost user=username dbname=database_name sslmode=disable password=your_password"

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return db
}
