package models

import "gorm.io/gorm"

type Role string

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     Role   `gorm:"default:basic_role" json:"role"`
}
