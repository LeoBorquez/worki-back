package routes

import (
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo, h *handler.Handler) {
	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)
	e.PATCH("/users/:id", h.UpdateUser)
}
