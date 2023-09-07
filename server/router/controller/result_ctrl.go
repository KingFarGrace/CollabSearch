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
				if resp := service.SetSearchingPhrase(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		resultGroup.POST(
			"/search",
			middleware.ValidateSearchingJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.SearchingJSON)
				if !ok {
					context.JSON(http.StatusBadRequest, "Invalid request json.")
					return
				}
				if hints := service.GetSearchingHints(jsonObj); hints != nil {
					context.JSON(http.StatusOK, gin.H{"Msg": "Success.", "Hints": hints})
				} else {
					context.JSON(http.StatusUnprocessableEntity, gin.H{"Msg": "Failed to get searching hints."})
				}
			})
		resultGroup.POST(
			"/avg",
			middleware.JWTAuth(),
			middleware.ValidateResultIndexJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.ResultIndex)
				if !ok {
					context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
					return
				}
				avg := service.GetResultAvgScore(jsonObj)
				context.JSON(http.StatusOK, gin.H{"Msg": "Success.", "Avg": avg})
			})
		resultGroup.POST(
			"/note",
			middleware.JWTAuth(),
			middleware.ValidateNoteJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.Note)
				if !ok {
					context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
					return
				}
				if resp := service.SetNote(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
		resultGroup.POST(
			"/note/all",
			middleware.JWTAuth(),
			middleware.ValidateResultIndexJSON(),
			middleware.JWTInterceptor(),
			func(context *gin.Context) {
				jsonObj, ok := context.MustGet("jsonObj").(entity.ResultIndex)
				if !ok {
					context.JSON(http.StatusBadRequest, gin.H{"Msg": "Invalid request json."})
					return
				}
				if resp := service.GetNotes(jsonObj); resp.Success() {
					context.JSON(http.StatusOK, resp)
				} else {
					context.JSON(http.StatusUnprocessableEntity, resp)
				}
			})
	}
}
