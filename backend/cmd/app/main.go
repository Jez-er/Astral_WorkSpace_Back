package main

import (
	"Astral/internal/jwt/controllers"
	"Astral/internal/jwt/database"
	"Astral/internal/jwt/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func handlerFunc() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	api := r.Group("/api")
	{
		public := api.Group("/public")
		{

			public.POST("/login", controllers.Login)

			public.POST("/signup", controllers.Signup)
		}
		protected := api.Group("/protected").Use(middleware.AuthZ())
		{
			protected.GET("/profile", controllers.Profile)
		}
	}
	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	router := handlerFunc()
	router.Run("localhost:8080")
}
