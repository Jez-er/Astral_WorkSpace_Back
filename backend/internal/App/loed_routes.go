package routes

import (
	"github.com/gin-gonic/gin"
	"Astral/internal/jwt/routes" // Импортируем ваши маршруты
)

// LoadRoutes загружает маршруты из всех файлов в папке routes внутри internal и её поддиректорий
func LoadRoutes(router *gin.RouterGroup) {
	routes.RegisterPublicRoutes(router)
	routes.RegisterProtectedRoutes(router)
}
