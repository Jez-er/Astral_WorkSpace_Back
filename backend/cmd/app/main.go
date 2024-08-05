package main

import (
	routes "Astral/internal/App"
	"Astral/internal/App/database"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func handlerFunc() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	url := os.Getenv("FRONT_END_URL")

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{url},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == url
		},
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})

	api := r.Group("/api")
	routes.LoadRoutes(api) // Загрузка маршрутов из вашего нового маршрутизатора

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	// Загрузка .env файла
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := handlerFunc()
	router.Run(":3001")
}
