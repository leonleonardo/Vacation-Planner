package models

import (
	"gorm.io/gorm"
)

// introducing user type for database
type User struct {
	gorm.Model
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	Username  string `json: "username" gorm: "primaryKey"`
	Password  string `json: "password"`
}
