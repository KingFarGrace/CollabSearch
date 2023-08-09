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
	workspaceGroup.Use(middleware.JWTAuth())
	{
		// POST "/workspace"
		workspaceGroup.POST(
			"",
			middleware.ValidateWorkspaceJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj := context.MustGet("jsonObj").(entity.Workspace)
				if resp := service.CreateWorkspace(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		// GET "/workspace"
		workspaceGroup.GET(
			"",
			middleware.ValidateUidJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj := context.MustGet("jsonObj").(int64)
				if resp := service.GetWorkspaces(jsonObj, false); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		// PUT "/workspace/"
		workspaceGroup.PUT(
			"",
			middleware.ValidateWorkspaceJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj := context.MustGet("jsonObj").(entity.Workspace)
				if resp := service.UpdateWorkspace(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		// GET "/workspace/handler"
		workspaceGroup.GET(
			"/handler",
			middleware.ValidateUidJSON(),
			middleware.ValidateRegisterJSON(),
			func(context *gin.Context) {
				jsonObj := context.MustGet("jsonObj").(int64)
				if resp := service.GetWorkspaces(jsonObj, true); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		// GET "/workspace/:wid"
		workspaceGroup.GET(
			"/:wid",
			func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{"msg": "Enter a workspace!"})
			})
	}
}
