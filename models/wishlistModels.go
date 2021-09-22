package model

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	User       *User     `json:"user"`
	PropertyId uint      `json:"property_id"`
	Property   *Property `json:"property"`
}
