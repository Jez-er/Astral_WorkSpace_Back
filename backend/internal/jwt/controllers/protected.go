package controllers

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* проверка есть ли пользователь в базе */
func Profile(c *gin.Context) {

	var user models.User

	email, _ := c.Get("email")

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "User Not Found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Could Not Get User Profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)
}
