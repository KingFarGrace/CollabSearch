package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
)

// JWTAuth is not a filter or interceptor. It will pass the request to next
// function in every situation, and set a flag "jwtAuth" to indicate whether the
// request pass JWT Authentication for following functions.
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		prefix := "Bearer "
		jwt := context.GetHeader("Authorization")
		if len(jwt) < len(prefix) {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		jwt = jwt[len(prefix):]
		if token := util.JWTValidate(jwt, conf.Salt); token == nil {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		context.Set("jwtAuth", true)
		context.Next()
	}
}
