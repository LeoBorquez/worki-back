package model

import (
	"github.com/jinzhu/gorm"
)

// User table : user
type User struct {
	gorm.Model
	Name     string `gorm:"size:35"`
	LastName string `gorm:"size:35"`
	Email    string `gorm:"type:varchar(255);unique_index"`
	Phone    string `gorm:"type:varchar(10)"`
	Password string
	UserType string
	Token    string
}

// CreateUser struct
type CreateUser struct {
	Email    string `gorm:"type:varchar(255);unique_index"`
	Password string
}

// UpdateUser struct
type UpdateUser struct {
	Name     string
	LastName string
	Phone    string `gorm:"type:varchar(10)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

type handler struct {
	db map[string]*User
}
