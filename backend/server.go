package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo!")
	})

	e.GET("/gigs", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GET gigs")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
