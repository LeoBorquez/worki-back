package model

import "github.com/jinzhu/gorm"

// User table : user
type User struct {
	gorm.Model
	Name     string `sql:"name"`
	LastName string `sql:"last_name"`
	Email    string `sql:"email"`
	Phone    string `sql:"phone"`
	Password string `sql:"password"`
	Toke     string `sql:"token"`
}
