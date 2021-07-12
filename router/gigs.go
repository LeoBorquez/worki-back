package router

import (
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/labstack/echo"
)

func GigRoute(e *echo.Echo, h *handler.Handler) {
	e.POST("/gigs", h.CreateGig)
	e.GET("/feed", h.FetchGig)
}
