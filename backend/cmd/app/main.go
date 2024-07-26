package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"Astral/internal/jwt/database"
	"Astral/internal/App"
)

func handlerFunc() *gin.Engine {
	r := gin.Default()
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
