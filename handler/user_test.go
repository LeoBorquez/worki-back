package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeoBorquez/worki-back/config"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"Email":"jon@labstack.com", "Password":"testing"}`
	cfg      = config.LoadConfig()
	db       = config.SetupDB(cfg)
)

func TestCreateUser(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Handler{DB: db}

	// Assertions
	if assert.NoError(t, h.Signup(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}

}
