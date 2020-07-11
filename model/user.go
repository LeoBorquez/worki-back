package model

import (
	"time"

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
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string
}
