package handler

import (
	"fmt"
	"github.com/labstack/echo"
)

/*
type GigRequest struct {
	Tittle   string  `sql:"tittle"`
	Desc     string  `sql:"desc"`
	Image    string  `sql:"image"`
	Hourly   float64 `sql:"hourly,type:real"`
	Location string  `sql:"location"`
	IsActive bool    `sql:"is_activate"`
}
*/
func (h *Handler) CreateGig(c echo.Context) (err error){
	u := userIDFromToken(c)
	fmt.Println(u)
	return err
}
