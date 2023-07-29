package main

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.New()
	conf.InitLogger()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "Hello Gin!"})
	})
	var addr = "localhost:8080"
	err := router.Run(addr)
	if err != nil {
		var build strings.Builder
		build.WriteString("Err: ")
		build.WriteString(err.Error())
		util.ServerLogger(addr, "Err:", "Fatal")
		return
	}
}
