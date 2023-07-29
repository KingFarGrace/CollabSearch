package main

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/router"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	app := gin.New()
	conf.InitLogger()
	app.Use(cors.New(conf.CorsConf))
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
