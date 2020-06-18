package main

import (
	"fmt"
	"github.com/LeoBorquez/workiBack/handler"
	"github.com/LeoBorquez/workiBack/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"

	"github.com/labstack/echo"
)

func main() {

	// Start echo
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
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
	defer db.Close()
	// Handler
	h := &handler.Handler{DB: db}

	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)

	port := GetPort()
	fmt.Println("[-] Listening on ...", port)

	e.Logger.Fatal(e.Start(port))

}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
		fmt.Println("[-] No Port Enviroment variable detected. Setting to ", port)
	}
	return ":" + port
}
