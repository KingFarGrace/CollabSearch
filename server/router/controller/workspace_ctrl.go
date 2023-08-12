package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkspaceController struct{}

func (receiver WorkspaceController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.JWTAuth())
	// POST "/workspace"
	routerGroup.POST(
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
	routerGroup.GET(
		"",
		middleware.ValidateUidJSON(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.UidJSON)
			if resp := service.GetWorkspaces(jsonObj.Uid, false); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// GET "/workspace/handler"
	routerGroup.GET(
		"/handler",
		middleware.ValidateUidJSON(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.UidJSON)
			if resp := service.GetWorkspaces(jsonObj.Uid, true); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// PUT "/workspace"
	routerGroup.PUT(
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
	// GET "/workspace/:wid"
	// TODO: get workspace info.
	routerGroup.GET(
		"/:wid",
		func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "Enter a workspace!"})
		})
}
