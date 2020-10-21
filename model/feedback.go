package model

import (
	"github.com/jinzhu/gorm"
)

// Feedback is to review gig publisher
type Feedback struct {
	gorm.Model
	UserID uint `gorm:"FOREIGNKEY"`
	Rate   uint
}
