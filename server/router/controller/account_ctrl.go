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
	accountGroup := routerGroup.Group("/user")
	{
		// GET "/user" -- service.Login
		accountGroup.GET(
			"",
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
		// POST "/user" -- service.Register
		accountGroup.POST(
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
		// PUT "/user" -- service.UpdateUser
		accountGroup.PUT(
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
	}
}
