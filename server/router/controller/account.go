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
type TestController struct{}

func (receiver *AccountController) Register(routerGroup *gin.RouterGroup) {
	// path: "/user", restful
	accountGroup := routerGroup.Group("/user")
	{
		// GET "/user" -- service.Login
		// TODO: TEST
		accountGroup.Use(middleware.ValidateLoginJSON()).Use(middleware.JWTAuth()).GET("", func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.LoginJSON)
			jwtAuth := context.MustGet("jwtAuth").(bool)
			if jwtAuth {
				// User data has been cached in the front-end or they won't get a legal token.
				context.JSON(http.StatusOK, gin.H{"msg": "Authentication approved."})
				return
			}
			if resp, token := service.LoginWithoutToken(jsonObj); resp.Success() {
				bearer := util.Concat("Bearer ", token)
				context.Header("Authorization", bearer)
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnauthorized, resp)
			}
		})
		// POST "/user" -- service.Register
		accountGroup.Use(middleware.ValidateRegisterJSON()).POST("", func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.RegisterJSON)
			if resp := service.Register(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	}
}

func (receiver *TestController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello gin!")
	})
}
