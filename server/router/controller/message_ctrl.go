package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct{}

func (receiver *MessageController) Register(routerGroup *gin.RouterGroup) {
	routerGroup.POST(
		"",
		middleware.JWTAuth(),
		middleware.ValidateMessage(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj, ok := context.MustGet("jsonObj").(entity.Message)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
				return
			}
			if resp := service.SendMessage(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		})
	routerGroup.GET(
		"/:receiver",
		middleware.JWTAuth(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			param := context.Param("receiver")
			receiver, ok := util.String2Int64(param)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
				return
			}
			if resp := service.GetMessages(receiver); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		},
	)
	routerGroup.PUT(
		"",
		middleware.JWTAuth(),
		middleware.ValidateMessage(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj, ok := context.MustGet("jsonObj").(entity.Message)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
				return
			}
			if resp := service.ReadMessage(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		},
	)
	routerGroup.DELETE(
		"",
		middleware.JWTAuth(),
		middleware.ValidateMessage(),
		middleware.JWTInterceptor(),
		func(context *gin.Context) {
			jsonObj, ok := context.MustGet("jsonObj").(entity.Message)
			if !ok {
				context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
				return
			}
			if resp := service.RemoveMessage(jsonObj); resp.Success() {
				context.JSON(http.StatusOK, resp)
			} else {
				context.JSON(http.StatusUnprocessableEntity, resp)
			}
		},
	)
}
