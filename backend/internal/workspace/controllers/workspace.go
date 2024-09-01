package workspace_controller

import (
	"Astral/internal/App/database"
	workspace_models "Astral/internal/workspace/models"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type AddUserStruct struct {
	ID  uint   `json:"id"`
	Key string `json:"key"`
}

func CreateSpace(c *gin.Context) {
	var workspace workspace_models.Workspace
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
		"Space": workspace,
	})
}

func DeleteSpace(c *gin.Context) {
	id := c.Param("id") // ID рабочего пространства
	var workspace workspace_models.Workspace

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
	var workspace workspace_models.Workspace

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
		workspace_models.Workspace{
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
		"WorkSpace": workspace,
	})
}

// Все пространства
func GetWorkspaces(c *gin.Context) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	user := c.Param("user")

	cachedWorkspaces, err := database.RedisClient.Get("workspaces").Bytes()
	if err != nil {
		workspaces := fetchWorkspaces(c, user)
		cachedWorkspaces, err = json.Marshal(workspaces)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Error marshal workspaces",
				"Err":   err,
			})
			c.Abort()
		}

		err = database.RedisClient.Set("workspaces", cachedWorkspaces, 10*time.Second).Err()
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Error set workspaces redis",
				"Err":   err,
			})
			c.Abort()
		}

		c.JSON(200, gin.H{
			"WorkSpace": workspaces,
			"Message":   "postgres",
		})
		return
	}
	workspaces := []*workspace_models.Workspace{}

	err = json.Unmarshal(cachedWorkspaces, &workspaces)
	if err != nil {
		c.JSON(500, gin.H{
			"Err":     err,
			"Message": "Error unmarshal workspace",
		})
	}

	c.JSON(200, gin.H{
		"workspaces": workspaces,
		"Message":    "redis",
	})
}

func fetchWorkspaces(c *gin.Context, user string) []*workspace_models.Workspace {
	workspace := []*workspace_models.Workspace{}
	res := database.GlobalDB.Where("? IN (key, key2, key3)", user).Find(&workspace)

	if res.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Error get workspaces",
			"Err":   res.Error,
		})
		c.Abort()
	}
	return workspace
}

// 1 пространство
func GetWorkspace(c *gin.Context) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	id := c.Param("id") // ID рабочего пространства

	cachedWorkspaces, err := database.RedisClient.Get("workspace").Bytes()
	if err != nil {
		workspace := fetchWorkspace(c, id)
		cachedWorkspaces, err = json.Marshal(workspace)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Error marshal workspace",
				"Err":   err,
			})
			c.Abort()
		}

		err = database.RedisClient.Set("workspace", cachedWorkspaces, 10*time.Second).Err()
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Error set workspace redis",
				"Err":   err,
			})
			c.Abort()
		}

		c.JSON(200, gin.H{
			"WorkSpace": workspace,
			"Message":   "postgres",
		})
		return
	}
	var workspace workspace_models.Workspace

	err = json.Unmarshal(cachedWorkspaces, &workspace)
	if err != nil {
		c.JSON(500, gin.H{
			"Err":     err,
			"Message": "Error unmarshal workspace",
		})
	}

	c.JSON(200, gin.H{
		"WorkSpace": workspace,
		"Message":   "redis",
	})
}

func fetchWorkspace(c *gin.Context, id string) workspace_models.Workspace {
	var workspace workspace_models.Workspace
	res := database.GlobalDB.Find(&workspace, "id = ?", id)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"Error": "Error Get WorkSpaces",
			"Err":   res.Error,
		})
		c.Abort()
	}
	return workspace
}

func AddSecondUser(c *gin.Context) {
	var addUsers AddUserStruct
	var workspace workspace_models.Workspace

	// Привязка JSON из запроса к структуре
	if err := c.ShouldBindJSON(&addUsers); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}

	// Поиск рабочего пространства по ID и ключу key2
	res := database.GlobalDB.Where("id = ?", addUsers.ID).First(&workspace)
	if res.Error != nil {
		log.Println("Workspace not found:", res.Error)
		c.JSON(404, gin.H{
			"Error": "Workspace not found",
		})
		return
	}

	// Обновление рабочего пространства
	res = database.GlobalDB.Model(&workspace).Where("id = ?", addUsers.ID).Updates(workspace_models.Workspace{
		Key2: addUsers.Key,
	})
	if res.Error != nil {
		log.Println("Error updating workspace:", res.Error)
		c.JSON(500, gin.H{
			"Error": "Failed to update workspace",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "success",
	})
}

func AddThirdUser(c *gin.Context) {
	var addUsers AddUserStruct
	var workspace workspace_models.Workspace

	// Привязка JSON из запроса к структуре
	if err := c.ShouldBindJSON(&addUsers); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"Error": "Invalid Inputs",
		})
		return
	}

	// Поиск рабочего пространства по ID и ключу key2
	res := database.GlobalDB.Where("id = ?", addUsers.ID).First(&workspace)
	if res.Error != nil {
		log.Println("Workspace not found:", res.Error)
		c.JSON(404, gin.H{
			"Error": "Workspace not found",
		})
		return
	}

	// Обновление рабочего пространства
	res = database.GlobalDB.Model(&workspace).Where("id = ?", addUsers.ID).Updates(workspace_models.Workspace{
		Key3: addUsers.Key,
	})
	if res.Error != nil {
		log.Println("Error updating workspace:", res.Error)
		c.JSON(500, gin.H{
			"Error": "Failed to update workspace",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message": "success",
	})
}
