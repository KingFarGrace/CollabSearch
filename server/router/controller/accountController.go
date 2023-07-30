package router

import "github.com/gin-gonic/gin"

func AccountController(app *gin.Engine) {
	accountGroup := app.Group("/user")
	{
		accountGroup.GET("")
		accountGroup.POST("")
	}
}
