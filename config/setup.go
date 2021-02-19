package config

import (
	"fmt"

	"github.com/LeoBorquez/worki-back/model"
	"github.com/jinzhu/gorm"

	//Only for connection
	_ "github.com/lib/pq"
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

	// Create the Uri
	dbURI := fmt.Sprintf("host=%s port=%s test=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	if uri != "" {
		dbURI = uri
	}
	fmt.Println("[-] URI db", dbURI)

	// Database connection
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Drop tables for development
	if dev == true {
		fmt.Println("[-] Dropping tables")
		db.Debug().DropTableIfExists(model.Gig{}, model.User{}, model.Proposal{})
		db.Debug().AutoMigrate(model.User{}, model.Gig{}, model.Proposal{})
		fmt.Println("[-] Creating faking data")
		model.FakeGig(db)
	}
	// Migrate model
	db.Debug().AutoMigrate(model.User{}, model.Gig{})

	return db
}
