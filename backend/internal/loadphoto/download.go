package loadphoto

import (
	"Astral/internal/App/database"
	md "Astral/internal/jwt/models"

	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	var user md.User
	userID := GetUserId(c, user)

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(404, gin.H{"Message": "Error get file", "Error": err})
		return
	}

	uploadDir := "./image"
	ext := filepath.Ext(file.Filename)

	newFileName := userID + ext

	pathFile := uploadDir + file.Filename
	AddFilePath(c, pathFile, userID, user)
	dst := filepath.Join(uploadDir, newFileName)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(404, gin.H{"Message": "Error save file", "Error": err})
		return
	}

	c.JSON(200, gin.H{"Message": "success message", "File name": newFileName})
}

func GetUserId(c *gin.Context, user md.User) string {
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return "Invalid UserID"
	}

	return user.UserId
}

func AddFilePath(c *gin.Context, path string, userID string, user md.User) {
	res := database.GlobalDB.Model(&user).Where("id = ?", userID).Updates(
		md.User{
			Image: path,
		})
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Add File Path",
			"Err":   res.Error,
		})
		c.Abort()
	}
}
