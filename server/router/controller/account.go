package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct{}
type TestController struct{}

func (receiver *AccountController) Register(routerGroup *gin.RouterGroup) {
	accountGroup := routerGroup.Group("/user")
	{
		accountGroup.GET("", func(context *gin.Context) {
			//
			context.JSON(http.StatusOK, "login")
		})
		accountGroup.POST("", func(context *gin.Context) {
			context.JSON(http.StatusOK, "register")
		})
	}
}

func (receiver *TestController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Hello gin!")
	})
}
