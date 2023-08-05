package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
)

var salt = ""

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwt := context.GetHeader("Authorization")[len("Bearer "):]
		if jwt == "" {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		if token := util.JWTValidate(jwt, salt); token == nil {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		context.Set("jwtAuth", true)
		context.Next()
	}
}
