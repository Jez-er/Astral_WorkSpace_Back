package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"Astral/internal/jwt/models"
)

/* Глобальная переменная для работы с базой данных */
var GlobalDB *gorm.DB

/* Функция для подключения базы данных */
func InitDatabase() (err error) {
	localhost := "localhost"
	db := "db"
	user := "user"
	port := "5432"
	pass := "pass"
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable",
		localhost,
		user,
		db,
		port,
		pass,
	)
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = GlobalDB.AutoMigrate(&models.User{}) // Используем структуру из models
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
	return
}
