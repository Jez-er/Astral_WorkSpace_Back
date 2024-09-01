package routes

import (
	workspace_controller "Astral/internal/workspace/controllers"

	"github.com/gin-gonic/gin"
)

func WorkSpacePublicRoutes(r *gin.RouterGroup) {
	workspace := r.Group("/workspace")
	{
		workspace.POST("/space", workspace_controller.CreateSpace)
		workspace.DELETE("/space/:id", workspace_controller.DeleteSpace) // ID рабочего пространства
		workspace.PUT("/space/:id", workspace_controller.UpdateSpace) // ID рабочего пространства
		workspace.GET("/spaces/:user", workspace_controller.GetWorkspaces) // ID пользователя
		workspace.GET("/space/:id", workspace_controller.GetWorkspace) // ID рабочего пространства
		workspace.PUT("/secondUser/",workspace_controller.AddSecondUser) 
		workspace.PUT("/thirdUser/",workspace_controller.AddThirdUser) 
	}
}
