package router

import (
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/gofiber/fiber"
)

func UserRoute(app *fiber.App) {
	app.Get("/signup", handler.SignUp())
}
