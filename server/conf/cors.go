package conf

import (
	"github.com/gin-contrib/cors"
)

var CorsConf = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
	ExposeHeaders:    []string{"Content-Length", "Content-Type", "Cache-Control", "Set-Token"},
	AllowCredentials: true,
}
