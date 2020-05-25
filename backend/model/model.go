package model

import (
	"fmt"
	"log"

	pg "github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

// TestExport testing export
func TestExport() {
	fmt.Println("test export")
}

// CreateCon create the connection to th db
func CreateCon() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "worki",
	})

	if db == nil {
		log.Panic("Failed to connect to database.\n")
	}
	log.Printf("Connected succesfully")

	// migration
	CreateGigTable(db)

	defer db.Close()
}
