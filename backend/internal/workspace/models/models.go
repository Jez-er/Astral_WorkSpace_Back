package workspace_models

import (
	"gorm.io/gorm"
)

/* Структура рабочего простарнства */
type Workspace struct {
	gorm.Model
	ID          uint   `gorm:"primarykey"`
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LogoColor   string `json:"logoColor"`
}


func (w *Workspace) CreateWorkSpaceRecord(db *gorm.DB) error {
	result := db.Create(&w)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
