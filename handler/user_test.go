package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeoBorquez/worki-back/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*model.CreateUser{
		"user1": &model.CreateUser{Email: "leoborquez@mail.com", Password: "123123123"},
	}
	userJSON = `{"Email":"jon@labstack.com", "Password":"testing"}`
)

func TestCreateUser(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}

}
