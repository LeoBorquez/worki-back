package handler

import "github.com/jinzhu/gorm"

// returns a handle to the DB object
type Handler struct {
	DB *gorm.DB
}

const (
	Key = "secret"
)
