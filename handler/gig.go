package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// CreateGig create a new gig
func (h *Handler) CreateGig(c echo.Context) (err error) {
	u := userIDFromToken(c)
	fmt.Println(u)

	return c.JSON(http.StatusCreated, u)
}
