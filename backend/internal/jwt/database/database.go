package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*  Глобальная переменная для работы с базой данных */
type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}
var GlobalDB *gorm.DB

/* Функция для подключения базы данных */
func InitDatabase() (err error) {
	localhost := "localhost"
	db := "db"
	user := "user"
	pass := "pass"
	dsn := fmt.Sprintf("host=%s  user=%s dbname=%s password=%s sslmode=disable",
		localhost,
		user,
		db,
		pass ,
	)
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	GlobalDB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
	return
}
