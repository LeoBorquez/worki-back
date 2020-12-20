package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/jinzhu/gorm"
)

// SetupDB the database
func SetupDB(cfg *Config) *gorm.DB {

	username := cfg.NameDB
	fmt.Println(username)
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	uri := os.Getenv("DATABASE_URL")
	dev, err := strconv.ParseBool(os.Getenv("DEV"))
	if err != nil {
		fmt.Println("Define enviroment")
	}

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
	if dev == true {
		fmt.Println("[-] Dropping tables")
		db.Debug().DropTableIfExists(model.Gig{})
		db.Debug().AutoMigrate(model.User{}, model.Gig{}, model.Proposal{})
		fmt.Println("[-] Creating faking data")
		model.FakeGig(db)
	}
	// Migrate model
	db.Debug().AutoMigrate(model.User{}, model.Gig{})

	return db
}
