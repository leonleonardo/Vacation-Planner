package routes

import (
	"gorm.io/gorm"
)

type DBRouter struct {
	DB *gorm.DB
}

// Takes in database and creates object for it to be passed
func NewConnection(db *gorm.DB) DBRouter {
	return DBRouter{db}
}
