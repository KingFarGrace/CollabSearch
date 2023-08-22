package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	// GET "/workspace/:uid"
	routerGroup.GET(
		"/user/:uid",
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			param := context.Param("uid")
			uid, ok := util.String2Int64(param)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request param(s)."})
				return
			}
			if resp := service.GetWorkspaces(uid, false); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// GET "/workspace/handler/:uid"
	routerGroup.GET(
		"/handler/:uid",
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			param := context.Param("uid")
			uid, ok := util.String2Int64(param)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request param(s)."})
				return
			}
			if resp := service.GetWorkspaces(uid, true); resp.Success() {
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
	routerGroup.GET(
		"/:wid",
		func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"Msg": "Enter a workspace!"})
		})
	routerGroup.GET(
		"/:wid/participants",
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			param := context.Param("wid")
			wid, err := strconv.Atoi(param)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request param(s)."})
				return
			}
			if resp := service.GetWorkspaceParticipants(wid); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
}
