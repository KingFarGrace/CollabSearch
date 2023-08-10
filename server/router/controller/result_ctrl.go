package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultController struct{}

func (receiver *ResultController) Register(routerGroup *gin.RouterGroup) {
	resultGroup := routerGroup.Group("/result")
	{
		resultGroup.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Hello")
		})
	}
}
