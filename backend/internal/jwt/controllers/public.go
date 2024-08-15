package controllers

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/auth"
	"Astral/internal/jwt/models"
	ws "Astral/internal/workspace/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UsData struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
	Image string `json:"image"`
}

/* обновленный токен на возврат */
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

/* регистрация */
func Signup(c *gin.Context) {
	id := uuid.New().String()
	user := models.User{
		UserId:      id,
		WorkspaceID: 1,
		WorkSpace: ws.Workspace{
			Key:         id,
			Title:       "First Name",
			Description: "Test",
			LogoColor:   "#fffff",
		},
	}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs ",
		})
		c.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"Error": "Error Hashing Password",
		})
		c.Abort()
		return
	}
	err = user.CreateUserRecord(database.GlobalDB)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating User",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Register",
	})
}

/* вход */
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	result := database.GlobalDB.Where("email = ?", payload.Email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"Error": "Invalid User Credentials",
		})
		c.Abort()
		return
	}
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	signedRefreshToken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Signing Token",
		})
		c.Abort()
		return
	}
	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedRefreshToken,
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

	c.JSON(200, gin.H{"Tokens": tokenResponse, "UserData": UsData{
		Name:        user.Name,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		UserId:      user.UserId,
		Image: user.Image,
	}})
}

func RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "No refresh token found",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "AuthService",
		ExpirationMinutes: 1,
		ExpirationHours:   12,
	}

	claims, err := jwtWrapper.ValidationToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid refresh token",
		})
		c.Abort()
		return
	}

	// Check if the refresh token has expired
	if claims.ExpiresAt < time.Now().Local().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Refresh token expired",
		})
		c.Abort()
		return
	}

	// Refresh tokens
	signedToken, err := jwtWrapper.GenerateToken(claims.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error signing new token",
		})
		c.Abort()
		return
	}

	signedRefreshToken, err := jwtWrapper.RefreshToken(claims.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error signing new refresh token",
		})
		c.Abort()
		return
	}

	// Set the new refresh token in the cookie
	c.SetCookie(
		"refresh_token",
		signedRefreshToken,
		60*60*24*30,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(200, gin.H{
		"Token":        signedToken,
		"RefreshToken": signedRefreshToken,
	})
}

func GetUserInfo(c *gin.Context) {
	// Извлечение токена из куки
	tokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Не предоставлен куки 'refresh_token'"})
		c.Abort()
		return
	}

	// Проверка JWT токена
	jwtWrapper := auth.JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	email, err := jwtWrapper.ValidationToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Неверный токен"})
		c.Abort()
		return
	}

	// Запрос пользователя из базы данных
	var user models.User
	fmt.Println(email.Email)
	if result := database.GlobalDB.Where("email = ?", email.Email).First(&user); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Пользователь не найден"})
		c.Abort()
		return
	}

	// Формирование ответа
	c.JSON(http.StatusOK, gin.H{
		"UserID": user.UserId,
		"Name":        user.Name,
		"DisplayName": user.DisplayName,
		"Email":       user.Email,
		"Image": user.Image,
	})
}
