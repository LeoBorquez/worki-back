package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeoBorquez/workiBack/model"
	"github.com/labstack/echo"
)

var (
	mockDB = map[string]*model.CreateUser{
		"jon@mail.com": &model.CreateUser{"Jon Snow", "jon@labstack.com"},
	}
	userJSON = `{"name":"Jon Snow", "email:"jon@mail.com"}`
)

func TestCreateUser(t *testing.T) {
	//Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
}
