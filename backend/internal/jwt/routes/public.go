package routes

import (
	"Astral/internal/jwt/controllers"
	gl "Astral/internal/Oauth2"
	"github.com/gin-gonic/gin"
)

// RegisterPublicRoutes регистрирует публичные маршруты
func RegisterPublicRoutes(r *gin.RouterGroup) {
	public := r.Group("/public")
	{
		public.POST("/login", controllers.Login)
		public.POST("/signup", controllers.Signup)
		public.GET("/refresh", controllers.RefreshToken)
	}
	oauth := r.Group("oauth")
	{
		oauth.GET("/google", gl.GoogleOauthLogin)
		oauth.GET("/github", gl.GithubOauthLogin)
	}

	callback := r.Group("callback")
	{
		callback.GET("/google", gl.GoogleCallBack)
		callback.GET("/github", gl.GithubCallBack)
	}
}
