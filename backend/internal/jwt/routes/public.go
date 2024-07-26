package routes

import (
	gl "Astral/internal/google"
	"Astral/internal/jwt/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// RegisterPublicRoutes регистрирует публичные маршруты
func RegisterPublicRoutes(r *gin.RouterGroup) {
	public := r.Group("/public")
	{
		public.POST("/login", controllers.Login)
		public.POST("/signup", controllers.Signup)
	}
	googl := r.Group("/google")
	{
		/* err := godotenv.Load()
		if err != nil {
			log.Fatal(".env file failed to load!")
		} */
		clientID := "1081042866166-n8rpg34g7jkk1a8qd0vlj8dco4u9kc8b.apps.googleusercontent.com"
		clientSecret := "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"
		clientCallbackURL := "http://localhost:5000/auth/google/callback"

		if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
			log.Fatal("Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
		}

		goth.UseProviders(
			google.New(clientID, clientSecret, clientCallbackURL),
		)

		googl.GET("/", gl.Home)
		googl.GET("/auth/:provider", gl.SignInWithProvider)
		googl.GET("/auth/:provider/callback", gl.CallbackHandler)
		googl.GET("/success", gl.Success)

	}
}
