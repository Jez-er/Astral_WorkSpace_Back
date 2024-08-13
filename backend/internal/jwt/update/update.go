package update

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateEmail(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Invalid inputs",
			"err":   err,
		})
		return
	}
	res := database.GlobalDB.Model(&user).Where("id = ?", user.UserId).Updates(
		models.User{
			Email: user.Email,
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
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Invalid inputs",
			"err":   err,
		})
		return
	}
	res := database.GlobalDB.Model(&user).Where("id = ?", user.UserId).Updates(
		models.User{
			DisplayName: user.DisplayName,
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
