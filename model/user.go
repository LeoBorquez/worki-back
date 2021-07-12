package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User table : test
type User struct {
	gorm.Model
	UserType            UserType
	FirstName, LastName string `gorm:"size:35"`
	Email               string `gorm:"type:varchar(255);unique_index" json:"email"`
	Phone               string `gorm:"type:varchar(10)"`
	Birthday            *time.Time
	Password            string
	Bio                 string `gorm:"size:300"`
	GigsDone            int
	Token               string
	Feedbacks           []Feedback //type struct
	Proposals           []Proposal
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

type LeaveReview struct {
	UserID uint
	Review string `gorm:"size:500"`
}
