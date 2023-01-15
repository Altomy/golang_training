package models

import (
	"Preload/Database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	StudentID uint    `json:"student_id"`
	Fair      float64 `json:"fair"`
	Status    int     `json:"status"`
	Product   string  `json:"product"`
}

func (bi *Bill) Save(context *gin.Context) error {

	return Database.DB.Create(&bi).Error
}