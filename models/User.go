package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not_null" json:"first_name"`
	LastName  string `gorm:"not_null" json:"last_name"`
	Email     string `gorm:"not_null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
