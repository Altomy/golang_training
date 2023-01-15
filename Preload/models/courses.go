package models

import (
	"Preload/Database"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title    string    `json:"title"`
	Duration int       `json:"duration"`
	Students []Student `json:"students" gorm:"many2many:students_courses;"`
}

func (c *Course) Save() error {
	return Database.DB.Save(&c).Error
}
