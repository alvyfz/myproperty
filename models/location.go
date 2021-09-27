package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID     uint   `json:"id"`
	City   string `json:"city"`
	Region string `json:"region"`
}
