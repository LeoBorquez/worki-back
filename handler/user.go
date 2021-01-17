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
	// Bind the struct to the context
	uc := new(model.CreateUser)
	if err := c.Bind(uc); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Can't bind object"}
	}
	// Validate
	if uc.Email == "" || uc.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid email or password"}
	}

	// Save user
	u := model.User{Email: uc.Email, Password: uc.Password}

	if err := h.DB.Create(u).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Couldn't save user"}
	}

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
	if err := h.DB.Table("users").Find(u, "email = ? and password = ?", u.Email, u.Password).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID // add the id to the token
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Delete password in response

	return c.JSON(http.StatusOK, u)
}

// UpdateUser the user by id
func (h *Handler) UpdateUser(c echo.Context) (err error) {
	userID := userIDFromToken(c)

	// Bind
	update := &model.UpdateUser{}
	if err := c.Bind(update); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Can't bind object"}
	}

	// Model to update
	u := &model.User{}
	// Find the record by ID
	if err := h.DB.First(u, userID).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Record not found"}
	}

	// Validations
	if update.Name == "" || update.LastName == "" || update.Phone == "" || update.Email == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Fields cannot be empty"}
	}

	if err := h.DB.Model(u).Updates(update).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusConflict, Message: "The user cannot be updated"}
	}

	return c.JSON(http.StatusOK, u)
}

// Get user ID
func userIDFromToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return uint(claims["id"].(float64))
}
