package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "Astral/internal/jwt/database"
    "Astral/internal/App"
)

func handlerFunc() *gin.Engine {
    r := gin.Default()

    // Настройка CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:5173"
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
	if (err != nil) {
		log.Fatal(err)
	}
	
	router := handlerFunc()
	router.Run(":3001")
}
