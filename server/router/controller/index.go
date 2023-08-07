package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (receiver *IndexController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(context *gin.Context) {
		// Server healthy testing.
		// TODO: index.
		context.JSON(http.StatusOK, "Hello gin!")
	})
}
