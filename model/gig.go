package model

import (
	"github.com/jinzhu/gorm"
)

// Gig table name : gigs
type Gig struct {
	gorm.Model         // id, created, updated
	UserID      uint   `gorm:"FOREIGNKEY"` //change to uint32
	CategoryID  uint   `gorm:"FOREIGNKEY"`
	Tittle      string `gorm:"size:255"`
	Description string `sql:"description"`
	Image       string `sql:"image"`
	Rate        float64
	Location    string `sql:"location"`
	IsActive    bool   `sql:"is_activate"`
	Proposals   []Proposal
}
