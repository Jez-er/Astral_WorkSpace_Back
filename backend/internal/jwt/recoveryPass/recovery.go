package recoverypass

import (
	"Astral/internal/App/database"
	"Astral/internal/jwt/models"
	crypto "crypto/rand"
	"log"
	"math/big"
	"net/smtp"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ChangePass struct {
	CheckCode string `json:"checkCode"`
	Code      string `json:"code"`
	Email     string `json:"email"`
	NewPass   string `json:"newPass"`
}

func ChangePassword(c *gin.Context) {
	var changePass ChangePass
	var user models.User
	if err := c.ShouldBindJSON(&changePass); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}
	if changePass.CheckCode != changePass.Code {
		c.JSON(400, gin.H{"Error": "Invalid Check Code"})
		return
	}
	UpdatedPassword(c, user, changePass)
	c.JSON(200, gin.H{
		"Message": "Success change password",
	})

}

func SendEmail(mail, code string) error {
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	recipientEmail := mail
	subject := "Восстановление пароля"
	body := code
	// Настройки SMTP-сервера
	smtpHost := "smtp.list.ru"
	smtpPort := "587"

	// Форматирование сообщения
	message := []byte("Subject: " + subject + "\r\n" +
		"From: " + senderEmail + "\r\n" +
		"To: " + recipientEmail + "\r\n" +
		"\r\n" +
		body)

	// Аутентификация
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Отправка письма
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		return err
	}

	return nil
}

func SendCode(c *gin.Context) {
	var user ChangePass
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}
	code := GeneratePassword()
	err := SendEmail(user.Email, code)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Send Email",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "Send code",
		"Code":    code,
	})

}

func GeneratePassword() string {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	code := strconv.Itoa(int(safeNum.Int64()))
	return code
}

func UpdatedPassword(c *gin.Context, user models.User, changePass ChangePass) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(changePass.NewPass), 14)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Error HashPass",
			"Err":   err,
		})
		c.Abort()
	}
	res := database.GlobalDB.Model(&user).Where("email = ?", changePass.Email).Updates(
		models.User{
			Password: string(bytes),
		})
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Updated Pass",
			"Err":   res.Error,
		})
		c.Abort()
	}
}
