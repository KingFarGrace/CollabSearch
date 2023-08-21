package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultController struct{}

func (receiver *ResultController) Register(routerGroup *gin.RouterGroup) {
	resultGroup := routerGroup.Group("/result")
	resultGroup.Use(middleware.JWTAuth())
	{
		resultGroup.POST(
			"",
			middleware.ValidateResultJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.Result)
				if !ok {
					context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
					return
				}
				if resp := service.SetSearchingHistory(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		resultGroup.GET(
			"/search",
			middleware.ValidateSearchingJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.SearchingJSON)
				if !ok {
					context.JSON(http.StatusBadRequest, "Invalid request json.")
					return
				}
				if resp := service.GetSearchingAdvice(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
	}
}
