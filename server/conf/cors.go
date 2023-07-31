package conf

import (
	"github.com/gin-contrib/cors"
)

// CorsConf
// CORS configuration
// AllowOrigins - Any
// AllowMethods - Except PATCH
// ExposeHeaders - New field "Set-Token" to restore user token
// AllowCredentials - Allow
// Else default.
var CorsConf = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
	ExposeHeaders:    []string{"Content-Length", "Content-Type", "Cache-Control", "Set-Token"},
	AllowCredentials: true,
}
