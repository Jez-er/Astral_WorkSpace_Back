package common_routes

import (
	"Astral/internal/common/loadphoto"

	"github.com/gin-gonic/gin"
)

func CommonRoutes(r *gin.RouterGroup) {
	workspace := r.Group("/common")
	{
		workspace.POST("/photo", loadphoto.DownloadFile)
	}
}
