package main

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	mw "github.com/KingFarGrace/CollabSearch/server/middleware"
	"github.com/KingFarGrace/CollabSearch/server/router"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	app := gin.New()
	conf.InitLogger()
	conf.InitPersistenceLayer()
	app.Use(mw.InitCors())
	router.InitRouter(app)
	var addr = "localhost:8080"
	err := app.Run(addr)
	if err != nil {
		var build strings.Builder
		build.WriteString("Err: ")
		build.WriteString(err.Error())
		util.ServerLogger(addr, "Err:", "Fatal")
		return
	}
}
