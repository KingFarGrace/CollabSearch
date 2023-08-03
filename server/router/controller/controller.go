package router

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(router *gin.RouterGroup)
}

type controllers struct {
	AccountController
	TestController
}

var Controllers = new(controllers)
