package router

import "github.com/gin-gonic/gin"

type Controller interface {
	// Register allows a controller to register itself with a router group.
	Register(router *gin.RouterGroup)
}

type controllers struct {
	AccountController
	TestController
}

var Controllers = new(controllers)
