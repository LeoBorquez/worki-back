package model

import "github.com/jinzhu/gorm"

// Comment model belongs to Gigs
type Comment struct {
	gorm.Model
	GigID       uint   `gorm:"FOREINGKEY"`
	UserID      uint   `gorm:"FOREIGNKEY"`
	Description string `gorm:"size:255"`
	Like        uint
}
