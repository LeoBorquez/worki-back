package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LeoBorquez/workiBack/handler"
	"github.com/LeoBorquez/workiBack/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[-] Error loading .env file")
	}

	cors := os.Getenv("cors_url")
	fmt.Printf("[-] Value CORS %v\n", cors)

	// Start echo
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cors},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Connect to the database
	db := model.SetupDB()

	// Handler
	h := &handler.Handler{DB: db}

	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/gigs", h.CreateGig)
	e.POST("/feed", h.FetchGig)

	port := getPort()
	fmt.Println("[-] Listening on ...", port)

	e.Logger.Fatal(e.Start(port))

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
		fmt.Println("[-] No Port Enviroment variable detected. Setting to ", port)
	}
	return ":" + port
}
