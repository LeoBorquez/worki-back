package handler

import (
	"github.com/jinzhu/gorm"
)

// Handler return a DB object
type Handler struct {
	DB *gorm.DB
}

const (
	Key = "secret"
)
