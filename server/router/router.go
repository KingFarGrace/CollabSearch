package router

import "github.com/gin-gonic/gin"

// InitRouter
// TODO: modularization
func InitRouter(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello Gin!"})
	})
}
