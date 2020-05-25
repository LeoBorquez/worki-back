package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	app "github.com/LeoBorquez/worki/app"
	m "github.com/LeoBorquez/worki/model"
)

func main() {

	m.Test()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", app.Login)

	// Restricted group
	e.GET("/", app.Accessible)

	// Restricted group
	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &app.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", app.Restricted)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo!")
	})

	/*
		e.DELETE("/gigs/:id", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "DELETE gig "+c.Param("id"))
		})
	*/
	e.Logger.Fatal(e.Start(":1323"))

}
