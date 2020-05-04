package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	m "github.com/LeoBorquez/worki/backend/model"
)

func main() {

	m.TestExport()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Echo!")
	})

	e.GET("/gigs", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GET gigs")
	})
	e.PUT("/gigs", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "PUT gigs")
	})
	e.DELETE("/gigs/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "DELETE gig "+c.Param("id"))
	})

	e.Logger.Fatal(e.Start(":1323"))

}
