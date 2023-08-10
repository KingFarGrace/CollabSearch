package router

import (
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	ctrl "github.com/KingFarGrace/CollabSearch/server/router/controller"
	"github.com/gin-gonic/gin"
)

var controllers = ctrl.Controllers

// InitRouter to init Gin router module.
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.InitCors())
	routerGroup := router.Group("")
	{
		controllers.IndexController.Register(routerGroup)
		controllers.AccountController.Register(routerGroup)
		controllers.WorkspaceController.Register(routerGroup)
		controllers.ResultController.Register(routerGroup)
		controllers.NoteController.Register(routerGroup)
	}
	return router
}
