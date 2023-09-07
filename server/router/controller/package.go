package router

import "github.com/gin-gonic/gin"

type Controller interface {
	// Register allows a controller to register itself with a router group.
	Register(*gin.RouterGroup)
}

type controllers struct {
	IndexController
	AccountController
	WorkspaceController
	ResultController
	NoteController
	MessageController
}

var Controllers = new(controllers)
