package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/gin-gonic/gin"
)

func SetEngine() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("engine", conf.GetMySQLEngine())
		context.Next()
	}
}
