package router

import "github.com/gin-gonic/gin"

type AccountController gin.RouterGroup

func (this AccountController) Register(router *gin.Engine) {
	accountGroup := router.Group("/user")
	{
		accountGroup.GET("")
		accountGroup.POST("")
	}
}
