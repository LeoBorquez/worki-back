package model

import (
	"github.com/jinzhu/gorm"
)

// Proposal table
type Proposal struct {
	gorm.Model
	UserID      uint   `gorm:"FOREIGNKEY"`
	GigID       uint   `gorm:"FOREIGNKEY"`
	Description string `gorm:"size:500"`
	Rate        float64
	TimeFrame   int
	StatusID    string
}
