package update

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/auth"
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
	res := database.GlobalDB.Model(&user).Where("user_id = ?", email.UserId).Updates(
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

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(email.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedRefreshToken, err := jwtWrapper.RefreshToken(email.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}

	// Установка рефреш-токена в куки
	c.SetCookie(
		"refresh_token",    // имя куки
		signedRefreshToken, // значение куки
		60*60*24*30,        // срок действия куки в секундах (30 дней)
		"/",                // путь
		"localhost",        // домен
		false,              // secure (установить true, если используется HTTPS)
		true,               // httpOnly
	)

	c.JSON(200, gin.H{
		"token": signedToken,
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
	res := database.GlobalDB.Model(&user).Where("user_id = ?", displayName.UserId).Updates(
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
