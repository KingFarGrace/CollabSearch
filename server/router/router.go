package router

import (
	ctrl "github.com/KingFarGrace/CollabSearch/server/router/controller"
	"github.com/gin-gonic/gin"
)

// InitRouter
// TODO: modularization
func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello Gin!"})
	})
	noAuthGroup := router.Group("")
	accountController := ctrl.AccountController{}
	accountController.SetPath("user")
	accountController.Register(noAuthGroup)
	return router
}
