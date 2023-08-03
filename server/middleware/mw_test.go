package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestInitCors(t *testing.T) {
	router := gin.Default()
	router.Use(InitCors()).GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "ok")
	})
	err := router.Run(":8888")
	if err != nil {
		return
	}
}

func TestValidateRegisterJSON(t *testing.T) {
	router := gin.Default()
	router.Use(ValidateRegisterJSON()).POST("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Success.")
	})
	err := router.Run(":8888")
	if err != nil {
		return
	}
}
