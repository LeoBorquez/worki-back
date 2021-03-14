package model

import "github.com/jinzhu/gorm"

type Status struct {
	gorm.Model
	Action      string
	Description string `gorm:"size:100"`
}
