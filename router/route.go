package router

import (
	"github.com/LeoBorquez/worki-back/config"
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

var cfg = config.LoadConfig()

func SetupRoutes(app *fiber.App) {
	// Handler "receiver" attached to the function type
	//h := &handler.Handler{DB: dbConnection()}

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/signup", handler.SignUp)

}

func dbConnection() *gorm.DB {
	// Connect to the database
	db := config.SetupDB(cfg)

	return db
}
