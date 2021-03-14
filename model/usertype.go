package model

import "github.com/jinzhu/gorm"

type UserType struct {
	gorm.Model
	Type string
}
