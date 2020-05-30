package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Token JWT claims struct
type Token struct {
	GigID uint
	jwt.StandardClaims
}

// Gig table name : gigs
type Gig struct {
	gorm.Model
	ID        int       `sql:"id, pk"`
	Tittle    string    `sql:"tittle"`
	Desc      string    `sql:"desc"`
	Image     string    `sql:"image"`
	Hourly    float64   `sql:"hourly,type:real"`
	Location  string    `sql:"location"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_activate"`
}
