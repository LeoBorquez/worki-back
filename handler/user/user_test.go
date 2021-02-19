package user

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/LeoBorquez/worki-back/handler"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mock     sqlmock.Sqlmock
	userJSON = `{"Email":"leoborquez@mail.com","Password":"123455"}`
	db       *sql.DB
	err      error
)

func TestCreateUser(t *testing.T) {
	//Setup

	db, mock, err = sqlmock.New()
	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler.Handler{gdb}

	// Assertions
	if assert.NoError(t, h.) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}

}
