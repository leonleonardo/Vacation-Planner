package handler

import (
	"gorm.io/gorm"
)

type dbHandler struct {
	DB *gorm.DB
}

func NewConnection(db *gorm.DB) handler {
	return dbHandler{db}
}