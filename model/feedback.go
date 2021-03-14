package model

import (
	"github.com/jinzhu/gorm"
)

// Feedback is to review gig publisher
type Feedback struct {
	gorm.Model
	UserID      uint `gorm:"FOREIGNKEY"`
	OwnerID     uint
	GigID       uint `gorm:"FOREIGNKEY"`
	Score       uint
	Description string
	Likes       int
}
