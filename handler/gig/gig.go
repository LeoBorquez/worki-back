package gig

import (
	"net/http"
	"strconv"

	"github.com/LeoBorquez/worki-back/handler/user"
	"github.com/LeoBorquez/worki-back/model"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/labstack/echo"
)

// CreateGig create a new gig
func (h *Handler) CreateGig(c echo.Context) (err error) {

	db := h.DB
	userID := user.UserIDFromToken(c)

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

	return c.JSON(http.StatusCreated, g)
}

// FetchGig return the last gigs added
func (h *Handler) FetchGig(c echo.Context) (err error) {
	userID := user.userIDFromToken(c)
	if userID == 0 {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Login requested"}
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve gigs from databse
	gigs := []*model.Gig{}
	db := h.DB

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at desc"},
	}, &gigs)

	return c.JSON(http.StatusAccepted, paginator)
}
