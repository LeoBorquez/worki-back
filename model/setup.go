package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

// Setup the database
func SetupDB() *gorm.DB {
	// Load the .env file
	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	uri := os.Getenv("DATABASE_URL")

	// Create the Uri
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	if uri != "" {
		dbUri = uri
	}

	// Database connection
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	// Migrate model
	db.Debug().AutoMigrate(&User{}, &Gig{})

	return db
}
