package workspace

import (
	"Astral/internal/jwt/database"
	"Astral/internal/jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateSpace(c *gin.Context) {
	var workspace models.Workspace
	err := c.ShouldBindJSON(&workspace)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	err = workspace.CreateWorkSpaceRecord(database.GlobalDB)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"Error": "Error Creating WorkSpace",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"Message": "Sucessfully Created WorkSpace",
		"Space":   workspace,
	})
}

func DeleteSpace(c *gin.Context) {
	id := c.Param("id") // ID рабочего пространства
	var workspace models.Workspace

	res := database.GlobalDB.Where("id = ?", id).Delete(&workspace)
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Deleted WorkSpace",
			"Err":   res.Error,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Message": "Deleted WorkSpace",
	})
}

func UpdateSpace(c *gin.Context) {
	id := c.Param("id") // ID рабочего пространства
	var workspace models.Workspace

	err := c.ShouldBindJSON(&workspace)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		c.Abort()
		return
	}
	res := database.GlobalDB.Model(&workspace).Where("id = ?", id).Updates(
		models.Workspace{
			Title:       workspace.Title,
			Description: workspace.Description,
			LogoColor:   workspace.LogoColor,
		})
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Updated WorkSpace",
			"Err":   res.Error,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Message":   "Updated WorkSpace",
		"WorkSpace": workspace,
	})
}
 // Все пространства
func GetWorkspaces(c *gin.Context) {
	id := c.Param("user_id") // ID пользователя
	var workspace models.Workspace

	res := database.GlobalDB.Find(&workspace, "id = ?", id)
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Get WorkSpaces",
			"Err":   res.Error,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Message":   "Get WorkSpaces",
		"WorkSpace": workspace,
	})
}
 // 1 пространство
func GetWorkspace(c *gin.Context) {
	id := c.Param("user_id") // ID рабочего пространства
	var workspace models.Workspace

	res := database.GlobalDB.Find(&workspace, "id = ?", id)
	if res.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"Error": "Error Get WorkSpaces",
			"Err":   res.Error,
		})
		c.Abort()
	}
	c.JSON(200, gin.H{
		"Message":   "Get WorkSpaces",
		"WorkSpace": workspace,
	})
}
