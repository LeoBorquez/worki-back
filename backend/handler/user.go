package handler

import (
	"net/http"

	"github.com/LeoBorquez/worki/model"
	"github.com/labstack/echo"
)

type UserResquest struct {
	Name     string `sql:"name"`
	LastName string `sql:"last_name"`
	Email    string `sql:"email"`
	Phone    string `sql:"phone"`
	Pasword  string `sql:"password"`
	Toke     string `sql:"token"`
}

// Signup user handler
func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Pasword == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := h.DB
	defer db.Close()

	return nil
}
