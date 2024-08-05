package database

import (
	"Astral/internal/jwt/models"
	workspace_models "Astral/internal/workspace/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* Глобальная переменная для работы с базой данных */
var GlobalDB *gorm.DB

/* Функция для подключения базы данных */
func InitDatabase() (err error) {
	localhost := "localhost"
	db := os.Getenv("DB_DB")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	pass := os.Getenv("DB_PASS")
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
	err = GlobalDB.AutoMigrate(&models.User{},&workspace_models.Workspace{}) // Используем структуру из models
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
	return
}
