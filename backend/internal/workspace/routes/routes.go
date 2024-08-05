package routes

import (
	workspace_controller "Astral/internal/workspace/controllers"

	"github.com/gin-gonic/gin"
)

func WorkSpacePublicRoutes(r *gin.RouterGroup) {
	workspace := r.Group("/workspace")
	{
		workspace.POST("/createSpace", workspace_controller.CreateSpace)
		workspace.DELETE("/deleteSpace/:id", workspace_controller.DeleteSpace) // ID рабочего пространства
		workspace.PUT("/updateSpace/:id", workspace_controller.UpdateSpace) // ID рабочего пространства
		workspace.GET("getSpaces/:user", workspace_controller.GetWorkspaces) // ID пользователя
		workspace.GET("getSpace/:id", workspace_controller.GetWorkspace) // ID рабочего пространства
	}
}
