package handler

import (
	"net/http"
	"time"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

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
	//defer db.Close()

	return c.JSON(http.StatusCreated, u)
}

// Login user
func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB
	defer db.Close()
	if err = db.Table("users").Find(u, "email = ? and password = ?", u.Email, u.Password).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Delete password in response

	return c.JSON(http.StatusOK, u)
}

// Get user ID
func userIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return uint(claims["id"].(float64))
}
