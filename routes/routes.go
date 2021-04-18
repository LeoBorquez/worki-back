package routes

import (
	"github.com/LeoBorquez/worki-back/config"
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var cfg = config.LoadConfig()

func Routes(e *echo.Echo) {
	// Handler "receiver" attached to the function type
	h := &handler.Handler{DB: dbConnection()}

	e.POST("/signup", h.SignUp)
	e.POST("/login", h.Login)
	e.POST("/gigs", h.CreateGig)
	e.GET("/feed", h.FetchGig)
	e.PATCH("/users/:id", h.UpdateUser)

}

func dbConnection() *gorm.DB {
	// Connect to the database
	db := config.SetupDB(cfg)

	return db
}
