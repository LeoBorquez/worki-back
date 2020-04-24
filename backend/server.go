package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	"github.com/labstack/echo/v4"
)

func main() {

	connect()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}

func connect() {
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

	defer db.Close()
}
