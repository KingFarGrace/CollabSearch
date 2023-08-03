package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct{}
type TestController struct{}

func (receiver *AccountController) Register(routerGroup *gin.RouterGroup) {
	accountGroup := routerGroup.Group("/user")
	{
		accountGroup.GET("", func(context *gin.Context) {
			// TODO: Login
			context.JSON(http.StatusOK, "login")
		})
		accountGroup.Use(middleware.ValidateRegisterJSON()).POST("", func(context *gin.Context) {
			jsonObj := context.MustGet("jsonObj").(entity.RegisterJSON)
			if response := service.Register(jsonObj); response.Success() {
				context.JSON(http.StatusOK, response)
			} else {
				context.JSON(http.StatusUnprocessableEntity, response)
			}
		})
	}
}

func (receiver *TestController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello gin!")
	})
}
