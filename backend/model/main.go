package model

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

// CreateCon create the connection to th db
func CreateCon() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "worki",
	})

	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Connected succesfully")

	CreateGigTable(db)

	defer db.Close()
}
