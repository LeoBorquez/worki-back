package model

import (
	"github.com/jinzhu/gorm"
)

// Gig table name : gigs
type Gig struct {
	gorm.Model         // id, created, updated
	UserID      uint   `gorm:"FOREIGNKEY"` //change to uint32
	CategoryID  uint   `gorm:"FOREIGNKEY"`
	StatusID    uint   `gorm:"FOREIGNKEY"`
	Tittle      string `gorm:"size:50"`
	Description string `gorm:"size:500"`
	TimeFrame   int
	Image       string `sql:"image"`
	Rate        float64
	Location    string `sql:"location"`
	Proposals   []Proposal
}
