package model

import (
	"github.com/jinzhu/gorm"
)

// User table : user
type User struct {
	gorm.Model
	Name     string
	LastName string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Phone    string `gorm:"type:varchar(9)"`
	Password string
	Token    string
}

// CreateUser struct
type CreateUser struct {
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

// UpdateUser struct
type UpdateUser struct {
	Name     string
	LastName string
	Phone    string `gorm:"type:varchar(9)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}
