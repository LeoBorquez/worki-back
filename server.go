package main

import (
	"net/http"

	"github.com/LeoBorquez/worki-back/config"
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/LeoBorquez/worki-back/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var cfg = config.LoadConfig()

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	log.Fatal(app.Listen(loadPort()))
	// fmt.Printf("%v, %T\n", const, const) print value and type of const
	cors := cfg.Cors
	log.Printf("[-] Value CORS %v\n", cfg.Cors)

	// Start echo
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cors},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	routes.Routes(e)

	e.Logger.Fatal(e.Start(loadPort()))

}

func loadPort() string {
	port := cfg.Port
	if port == "" {
		port = "1323"
		log.Printf("[-] No Port Enviroment variable detected. Setting to ", port)
	}

	return ":" + port
}
