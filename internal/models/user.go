package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ExternalId string
	Role       string
	Email      string
	Password   string
}
