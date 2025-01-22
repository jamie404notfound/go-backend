package models

import "gorm.io/gorm"

type LoginInput struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
