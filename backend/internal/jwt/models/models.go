package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/* Структура юзера для базы данных */
type User struct {
	gorm.Model
	UserId      string `gorm:"primarykey"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required" gorm:"unique"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	HowDid      string `json:"howDid" binding:"required"`
	WorkSpace   Workspace
}

/* Структура рабочего простарнства */
type Workspace struct {
	gorm.Model
	ID          string `gorm:"primarykey"`
	UserID      string
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	LogoColor   string `json:"logoColor" binding:"required"`
}

/* Создает таблицу по юзеру в базе данных */
func (u *User) CreateUserRecord(db *gorm.DB) error {
	result := db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/* Хеширование пароля */
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

/* Сравнивает шифрованный и обыный пароль */
func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (w *Workspace) CreateWorkSpaceRecord(db *gorm.DB) error {
	result := db.Create(&w)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
