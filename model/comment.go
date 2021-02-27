package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	GigID uint `gorm:"FOREINGKEY"`
	UserID  uint   `gorm:"FOREIGNKEY"`
	Comment string `gorm:"size:255"`
	Likes   int
}
