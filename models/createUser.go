package models

import "gorm.io/gorm"

type Body struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}
