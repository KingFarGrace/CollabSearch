package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitCors activate CORS configuration.
// @See conf.CorsConf
func InitCors() gin.HandlerFunc {
	return cors.New(conf.CorsConf)
}
