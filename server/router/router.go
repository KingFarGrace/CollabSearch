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
	indexGroup := router.Group("/")
	{
		controllers.IndexController.Register(indexGroup)
	}
	accountGroup := router.Group("/user")
	{
		controllers.AccountController.Register(accountGroup)
	}
	workspaceGroup := router.Group("/workspace")
	{
		controllers.WorkspaceController.Register(workspaceGroup)
		controllers.ResultController.Register(workspaceGroup)
		controllers.NoteController.Register(workspaceGroup)
	}
	messageGroup := router.Group("/msg")
	{
		controllers.MessageController.Register(messageGroup)
	}
	return router
}
