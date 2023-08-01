package router

import "github.com/gin-gonic/gin"

// InitRouter
// TODO: modularization
func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello Gin!"})
	})
	return router
}
