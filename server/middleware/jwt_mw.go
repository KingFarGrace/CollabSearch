package middleware

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuth is not a filter or interceptor. It will pass the request to following
// functions in every situation, and set a flag "jwtAuth" to indicate whether the
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
		token := util.JWTValidate(jwt, conf.Salt)
		if token == nil {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		claims := entity.TokenClaims{}
		err := sonic.Unmarshal(token.Claims(), &claims)
		if err != nil {
			context.Set("jwtAuth", false)
			context.Next()
			return
		}
		context.Set("claims", claims)
		context.Set("jwtAuth", true)
		context.Next()
	}
}

// JWTInterceptor will abort the request without JWT authentication.
// In most cases, this middleware should only be used after JWTAuth.
func JWTInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtAuth := context.MustGet("jwtAuth").(bool)
		if !jwtAuth {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Authentication required."})
			return
		}
		context.Next()
	}
}
