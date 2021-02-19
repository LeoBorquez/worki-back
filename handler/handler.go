package handler

import (
	"github.com/jinzhu/gorm"
)

// Handler retrun a DB object
type Handler struct {
	DB *gorm.DB
}

const (
	Key = "secret"
)
