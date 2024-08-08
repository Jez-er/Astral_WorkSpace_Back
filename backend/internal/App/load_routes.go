package routes

import (
	common_routes "Astral/internal/common/routes"
	"Astral/internal/jwt/routes"
	workspace "Astral/internal/workspace/routes" // Импортируем ваши маршруты

	"github.com/gin-gonic/gin"
)

// LoadRoutes загружает маршруты из всех файлов в папке routes внутри internal и её поддиректорий
func LoadRoutes(router *gin.RouterGroup) {
	/* Авторизация | Регистрация */
	routes.RegisterPublicRoutes(router)
	routes.RegisterProtectedRoutes(router)
	/* Рабочие пространство */
	workspace.WorkSpacePublicRoutes(router)
	/* Коммон */
	common_routes.CommonRoutes(router)
}
