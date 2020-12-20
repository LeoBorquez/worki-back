package config

import (
	"fmt"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/jinzhu/gorm"
)

// SetupDB the database
func SetupDB(cfg *Config) *gorm.DB {

	username := cfg.UserDB
	password := cfg.PassDB
	dbName := cfg.NameDB
	dbHost := cfg.HostDB
	dbPort := cfg.PortDB
	uri := cfg.URI
	dev := cfg.Dev
	if dev == false {
		fmt.Println("Running on Production Enviroment")
	}

	fmt.Printf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

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