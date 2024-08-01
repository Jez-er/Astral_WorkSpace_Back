package workspace

import (
	"github.com/gin-gonic/gin"
)

func WorkSpacePublicRoutes(r *gin.RouterGroup) {
	workspace := r.Group("/workspace")
	{
		workspace.POST("/createSpace", CreateSpace)
		workspace.DELETE("/deleteSpace/:id", DeleteSpace) // ID рабочего пространства
		workspace.PUT("/updateSpace/:id", UpdateSpace) // ID рабочего пространства
		workspace.GET("getSpaces/:user_id", GetWorkspaces) // ID пользователя
		workspace.GET("getSpace/:id", GetWorkspace) // ID рабочего пространства
	}
}
