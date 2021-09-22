package model

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ID          uint       `json:"id"`
	UserID      uint       `json:"user_id"`
	User        *User      `json:"user,omitempty"`
	DeveloperId uint       `json:"developer_id"`
	Developer   *Developer `json:"developer,omitempty"`
}
