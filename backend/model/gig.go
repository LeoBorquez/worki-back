package model

import (
	"github.com/jinzhu/gorm"
)

// Gig table name : gigs
type Gig struct {
	gorm.Model         // id, created, updated
	UserID     uint    `sql:"user_id"`
	Tittle     string  `sql:"tittle"`
	Desc       string  `sql:"desc"`
	Image      string  `sql:"image"`
	Hourly     float64 `sql:"hourly,type:real"`
	Location   string  `sql:"location"`
	IsActive   bool    `sql:"is_activate"`
}
