package router

import (
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	ctrl "github.com/KingFarGrace/CollabSearch/server/router/controller"
	"github.com/gin-gonic/gin"
)

var controllers = ctrl.Controllers

// InitRouter
// TODO: modularization
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.InitCors())
	noAuthGroup := router.Group("")
	{
		controllers.TestController.Register(noAuthGroup)
		controllers.AccountController.Register(noAuthGroup)
	}
	return router
}
