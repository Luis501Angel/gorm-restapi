package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not_null"`
	LastName  string `gorm:"not_null"`
	Email     string `gorm:"not_null;unique_index"`
	Tasks     []Task
}
