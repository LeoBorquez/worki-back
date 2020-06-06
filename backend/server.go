package main

import (
	"fmt"
	"github.com/LeoBorquez/worki/handler"
	"github.com/LeoBorquez/worki/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"

	"github.com/labstack/echo"
)

func main() {

	// Start echo
	e := echo.New()

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

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Migration &User{}
	db.Debug().AutoMigrate(&model.User{}, &model.Gig{})

	h := &handler.Handler{DB: db}

	e.POST("/signup", h.Signup)

	/*
		e.DELETE("/gigs/:id", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "DELETE gig "+c.Param("id"))
		})
	*/
	e.Logger.Fatal(e.Start(":1323"))

}
