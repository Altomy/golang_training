package models

import (
	"Preload/Database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Email   string   `json:"email"`
	Bills   []Bill   `json:"bill,omitempty"`
	Profile Profile  `json:"profile,omitempty"`
	Courses []Course `json:"courses" gorm:"many2many:students_courses;"`
}

func (st *Student) Save(context *gin.Context) error {
	return Database.DB.Create(&st).Error
}
