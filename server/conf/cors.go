package conf

import (
	"github.com/gin-contrib/cors"
)

var corsConfig = GetConfig().Cors

// CorsConf
// CORS configuration
// AllowOrigins - Any
// AllowMethods - Except PATCH
// AllowCredentials - Allow
// Else default.
var CorsConf = cors.Config{
	AllowOrigins:     corsConfig.AllowOrigins,
	AllowMethods:     corsConfig.AllowMethods,
	AllowHeaders:     corsConfig.AllowHeaders,
	ExposeHeaders:    corsConfig.ExposeHeaders,
	AllowCredentials: corsConfig.AllowCredentials,
}
