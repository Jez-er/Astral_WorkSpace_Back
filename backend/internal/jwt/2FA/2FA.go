package fa

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
)

type FAApply struct {
	CheckCode string `json:"checkCode"`
	Code      string `json:"code"`
	Email     string `json:"email"`
}

func FA(c *gin.Context) {
	var fa FAApply
	var user models.User
	if err := c.ShouldBindJSON(&fa); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}
	if fa.CheckCode != fa.Code {
		c.JSON(400, gin.H{"Error": "Invalid Check Code"})
		return
	}
	UpdatedFA(c, user, fa)
	c.JSON(200, gin.H{
		"Message": "Success 2FA",
	})
}

func SendEmail(mail, code string) error {
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	recipientEmail := mail
	subject := "Подтвердить почту"
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
	var fa FAApply
	if err := c.ShouldBindJSON(&fa); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}
	code := GenerateCode()
	err := SendEmail(fa.Email, code)
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

func GenerateCode() string {
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	code := strconv.Itoa(int(safeNum.Int64()))
	return code
}

func UpdatedFA(c *gin.Context, user models.User, fa FAApply) {
	res := database.GlobalDB.Model(&user).Where("email = ?", fa.Email).Updates(
		models.User{
			FA: true,
		})
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Updated 2FA",
			"Err":   res.Error,
		})
		c.Abort()
	}
}

func Change2FA(c *gin.Context) {
	var fa FAApply
	var user models.User
	if err := c.ShouldBindJSON(&fa); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}

	res := database.GlobalDB.Model(&user).Where("email = ? AND fa = ?", fa.Email, true).Updates(
		models.User{
			FA: false,
		})
	if res.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Error Updated 2FA",
			"Err":   res.Error,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Message": "success",
		"user":    user,
		"2FA":     user.FA,
	})
}
