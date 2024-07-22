package models

import (
	"Astral/internal/jwt/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/* Cтруктура юзера для базы данных */

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

/* Создает таблицу по юзеру в базе данных */
func (u *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&u)
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
