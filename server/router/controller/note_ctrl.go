package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NoteController struct{}

func (receiver *NoteController) Register(routerGroup *gin.RouterGroup) {
	noteGroup := routerGroup.Group("/note")
	{
		noteGroup.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Hello note.")
		})
	}
}
