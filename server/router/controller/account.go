package router

import "github.com/gin-gonic/gin"

type AccountController struct {
	RelativePath string
}

func (receiver *AccountController) Register(routerGroup *gin.RouterGroup) {
	accountGroup := routerGroup.Group(receiver.RelativePath)
	{
		accountGroup.GET("")
		accountGroup.POST("")
	}
}

func (receiver *AccountController) SetPath(path string) {
	receiver.RelativePath = path
}
