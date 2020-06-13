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
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := h.DB
	if err := db.Create(u); err == nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Couldn't save user"}
	}
	defer db.Close()

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (error error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}
