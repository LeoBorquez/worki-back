package main

import (
	"fmt"
	"github.com/LeoBorquez/workiBack/handler"
	"github.com/LeoBorquez/workiBack/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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

	// Database connection
	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	uri := os.Getenv("DATABASE_URL")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	if uri != "" {
		dbURI = uri
	}
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migration &User{}
	db.Debug().AutoMigrate(&model.User{}, &model.Gig{})

	// Handler
	h := &handler.Handler{DB: db}

	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":1323"
	}

	fmt.Printf(port)
	e.Logger.Fatal(e.Start(port))

}
