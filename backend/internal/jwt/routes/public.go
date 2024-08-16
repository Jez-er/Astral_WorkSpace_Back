package routes

import (
	gl "Astral/internal/Oauth2"
	"Astral/internal/jwt/controllers"
	up "Astral/internal/jwt/recoveryPass"
	upd "Astral/internal/jwt/update"

	"github.com/gin-gonic/gin"
)

// RegisterPublicRoutes регистрирует публичные маршруты
func RegisterPublicRoutes(r *gin.RouterGroup) {
	public := r.Group("/public")
	{
		public.POST("/login", controllers.Login)
		public.POST("/signup", controllers.Signup)
		public.GET("/refresh", controllers.RefreshToken)
		public.GET("/data", controllers.GetUserInfo)
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
	forget := r.Group("forget")
	{
		forget.POST("/sendCode", up.SendCode)
		forget.PUT("/updatePass", up.ChangePassword)
	}
	update := r.Group("/update")
	{
		update.PUT("/email", upd.UpdateEmail)
		update.PUT("/displayName", upd.UpdateDisplayName)
	}
}
