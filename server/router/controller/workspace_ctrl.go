package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkspaceController struct{}

func (receiver WorkspaceController) Register(router *gin.RouterGroup) {
	workspaceGroup := router.Group("/workspace")
	{
		// POST "/workspace"
		workspaceGroup.POST("", middleware.JWTAuth(), middleware.ValidateWorkspaceJSON(), func(context *gin.Context) {
			jwtAuth := context.MustGet("jwtAuth").(bool)
			if !jwtAuth {
				context.JSON(http.StatusUnauthorized, gin.H{"msg": "Authentication required."})
				return
			}
			jsonObj := context.MustGet("jsonObj").(entity.Workspace)
			if resp := service.CreateWorkspace(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
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
			context.JSON(http.StatusOK, gin.H{"msg": "UpdateUser info of workspace!"})
		})
	}
}
