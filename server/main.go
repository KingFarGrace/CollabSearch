package main

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/router"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"strings"
)

func main() {
	app := router.InitRouter()
	conf.InitLogger()
	conf.InitPersistenceLayer()
	conf.InitCachingLayer()
	var addr = "localhost:8080"
	serverName := util.Concat("Server(", addr, ")")
	err := app.Run(addr)
	if err != nil {
		var build strings.Builder
		build.WriteString("Err: ")
		build.WriteString(err.Error())
		util.LaunchLogger(serverName, false)
		util.FatalLogger(err, "app.Run")
		return
	}
}
