package model

import (
	"github.com/jinzhu/gorm"
)

// Gig table name : gigs
type Gig struct {
	gorm.Model          // id, created, updated
	UserID      uint    `gorm:"FOREIGNKEY"`
	Tittle      string  `gorm:"size:255"`
	Description string  `sql:"description"`
	Image       string  `sql:"image"`
	Pay         float64 `sql:"hourly,type:real"`
	Location    string  `sql:"location"`
	IsActive    bool    `sql:"is_activate"`
}
