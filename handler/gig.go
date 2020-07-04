package handler

import (
	"net/http"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/labstack/echo"
)

// CreateGig create a new gig
func (h *Handler) CreateGig(c echo.Context) (err error) {

	db := h.DB
	userID := userIDFromToken(c)

	g := &model.Gig{}
	u := &model.User{}
	g.UserID = userID

	// binding gig
	if err = c.Bind(g); err != nil {
		return
	}

	// validation
	if g.Tittle == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid tittle field"}
	}

	if db.Table("users").Find(u, userID).RecordNotFound() {
		return echo.ErrNotFound
	}

	if err := db.Create(g); err == nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Couldn't save gig"}
	}

	db.Create(g)
	defer db.Close()
	return c.JSON(http.StatusCreated, g)
}
