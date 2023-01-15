package models

import (
	"Preload/Database"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	StudentID uint   `json:"student_id"`
	Image     string `json:"image"`
	Address   string `json:"address"`
	City      string `json:"city"`
}

func (p *Profile) Save() error {
	return Database.DB.Save(&p).Error
}
