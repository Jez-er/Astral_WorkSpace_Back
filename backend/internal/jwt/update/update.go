package update

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`
}

type DisplayName struct {
	DisplayName string `json:"displayName"`
	UserId      string `json:"userId"`
}

func UpdateEmail(c *gin.Context) {
	var email Email
	var user models.User
	if err := c.ShouldBindJSON(&email); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Invalid inputs",
			"err":   err,
		})
		return
	}
	res := database.GlobalDB.Model(&user).Where("id = ?", email.UserId).Updates(
		models.User{
			Email: email.Email,
		})
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(500, gin.H{
			"Error": "Invalid update",
			"err":   res.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "success",
		"User":    user,
	})
}

func UpdateDisplayName(c *gin.Context) {
	var displayName DisplayName
	var user models.User
	if err := c.ShouldBindJSON(&displayName); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Invalid inputs",
			"err":   err,
		})
		return
	}
	res := database.GlobalDB.Model(&user).Where("id = ?", displayName.UserId).Updates(
		models.User{
			DisplayName: displayName.DisplayName,
		})
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(500, gin.H{
			"Error": "Invalid update",
			"err":   res.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "success",
		"User":    user,
	})
}
