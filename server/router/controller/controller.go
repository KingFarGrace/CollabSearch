package router

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(router *gin.RouterGroup)
	SetPath(path string)
}
