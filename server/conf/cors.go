package conf

import (
	"github.com/gin-contrib/cors"
)

var corsConfig = GetConfig().Cors

// CorsConf
// CORS configuration
// AllowOrigins - Any
// AllowMethods - Except PATCH
// ExposeHeaders - New field "Set-Token" to restore user token
// AllowCredentials - Allow
// Else default.
var CorsConf = cors.Config{
	AllowOrigins:     corsConfig.AllowOrigins,
	AllowMethods:     corsConfig.AllowMethods,
	ExposeHeaders:    corsConfig.ExposeHeaders,
	AllowCredentials: corsConfig.AllowCredentials,
}
