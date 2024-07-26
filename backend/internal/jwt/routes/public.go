package routes

import (
	"Astral/internal/jwt/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterPublicRoutes регистрирует публичные маршруты
func RegisterPublicRoutes(r *gin.RouterGroup) {
	public := r.Group("/public")
	{
		public.POST("/login", controllers.Login)
		public.POST("/signup", controllers.Signup)
	}
}
