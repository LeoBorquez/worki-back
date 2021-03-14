package model

import (
	"github.com/jinzhu/gorm"
)

// Proposal table
type Proposal struct {
	gorm.Model
	UserID      uint   `gorm:"FOREINGKEY"`
	GigID       uint   `gorm:"FOREINGKEY"`
	StatusID    uint   `gorm:"FOREINGKEY"`
	Description string `gorm:"size:500"`
	Rate        float64
	TimeFrame   int
}
