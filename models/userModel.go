package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	LastLogin time.Time `json:"last_login"`
}
