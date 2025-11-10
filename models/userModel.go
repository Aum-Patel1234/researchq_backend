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

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UpdateUserRequest struct {
	Name      *string    `json:"name,omitempty"`
	Email     *string    `json:"email,omitempty" binding:"omitempty,email"`
	Password  *string    `json:"password,omitempty" binding:"omitempty,min=8"`
	LastLogin *time.Time `json:"last_login,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
