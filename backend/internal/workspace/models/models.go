package workspace_models

import (
	"gorm.io/gorm"
)

/* Структура рабочего простарнства */
type Workspace struct {
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	Key         string `json:"key" binding:"required"`
	Key2        string `json:"key2"`
	Key3        string `json:"key3"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	LogoColor   string `json:"logoColor" binding:"required"`
}

func (w *Workspace) CreateWorkSpaceRecord(db *gorm.DB) error {
	result := db.Create(&w)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
