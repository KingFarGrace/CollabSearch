package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct{}

func (receiver *AccountController) Register(routerGroup *gin.RouterGroup) {
	// GET "/user/check"
	routerGroup.GET(
		"/check",
		middleware.JWTAuth(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			claims, exists := context.Get("claims")
			if !exists {
				context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Msg": "Failed to get user claims."})
				return
			}
			if resp := service.CheckStatus(claims.(entity.TokenClaims)); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnauthorized, resp)
			}
		})
	// POST "/user/login"
	routerGroup.POST(
		"/login",
		middleware.ValidateLoginJSON(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.LoginJSON)
			if resp, token := service.LoginWithoutToken(jsonObj); resp.Success() {
				bearer := util.Concat("Bearer ", token)
				context.Header("Authorization", bearer)
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnauthorized, resp)
			}
		})
	// GET "/user/:key
	routerGroup.GET(
		"/:key",
		middleware.JWTAuth(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			key := context.Param("key")
			if resp := service.SearchUser(key); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// POST "/user"
	routerGroup.POST(
		"",
		middleware.ValidateRegisterJSON(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.RegisterJSON)
			if resp := service.Register(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// PUT "/user"
	routerGroup.PUT(
		"",
		middleware.JWTAuth(),
		middleware.ValidateUpdateJSON(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.User)
			if resp := service.UpdateUser(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	// POST "/user/workspace"
	routerGroup.POST(
		"/workspace",
		middleware.JWTAuth(),
		middleware.ValidateUserWorkspaceJSON(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.UserWorkspace)
			if resp := service.JoinWorkspace(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
}
