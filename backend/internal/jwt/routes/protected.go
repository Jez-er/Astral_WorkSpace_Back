package routes

import (
	"Astral/internal/jwt/controllers"
	"Astral/internal/jwt/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterProtectedRoutes регистрирует защищенные маршруты
func RegisterProtectedRoutes(r *gin.RouterGroup) {
	protected := r.Group("/protected").Use(middleware.AuthZ())
	{
		protected.GET("/profile", controllers.Profile)
	}
}
