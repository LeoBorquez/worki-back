package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// SetupDB the database
func SetupDB() *gorm.DB {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	uri := os.Getenv("DATABASE_URL")
	dev := os.Getenv("dev")

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
