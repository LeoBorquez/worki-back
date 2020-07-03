package handler

import (
	"net/http"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/labstack/echo"
)

// CreateGig create a new gig
func (h *Handler) CreateGig(c echo.Context) (err error) {
	u := userIDFromToken(c)
	g := &model.Gig{}

	g.UserID = u
	g.Tittle = "test gig"

	h.DB.Create(g)
	return c.JSON(http.StatusCreated, u)
}
