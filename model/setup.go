package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// SetupDB the database
func SetupDB() *gorm.DB {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	uri := os.Getenv("DATABASE_URL")
	dev := os.Getenv("DEV")

	// Create the Uri
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	if uri != "" {
		dbURI = uri
	}

	// Database connection
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Drop tables for development
	if dev == "y" {
		db.Debug().DropTableIfExists(&Gig{})
		db.Debug().AutoMigrate(&User{}, &Gig{})
		FakeGig(db)
	}
	// Migrate model
	db.Debug().AutoMigrate(&User{}, &Gig{})

	return db
}
