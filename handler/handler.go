package handler

import (
	"os/user"

	"github.com/jinzhu/gorm"
)

// Handler retrun a DB object
type Handler struct {
	DB *gorm.DB
}

type User struct {
	*user.User
}

const (
	Key = "secret"
)
