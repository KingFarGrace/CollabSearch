package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkspaceController struct{}

func (receiver WorkspaceController) Register(router *gin.RouterGroup) {
	workspaceGroup := router.Group("/workspace")
	{
		// POST "/workspace"
		workspaceGroup.POST("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Create new workspace!"})
		})
		// GET "/workspace"
		workspaceGroup.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Get all workspaces you joined!"})
		})
		// GET "/workspace/handler"
		workspaceGroup.GET("/handler", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Get all workspaces created by you!"})
		})
		// GET "/workspace/:wid"
		workspaceGroup.GET("/:wid", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Enter a workspace!"})
		})
		// PUT "/workspace/:wid"
		workspaceGroup.PUT("/:wid", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Update info of workspace!"})
		})
	}
}
