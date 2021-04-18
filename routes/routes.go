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

	UserRoute(e, h)
	


}

func dbConnection() *gorm.DB {
	// Connect to the database
	db := config.SetupDB(cfg)

	return db
}
