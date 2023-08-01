package conf

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

var mysqlConf = GetConfig().GetXormConf()
var mysqlEngine *xorm.Engine
var redisConf = GetConfig().GetRedisConf()
var redisClient *redis.Client

// InitPersistenceLayer to connect MySQL database.
func InitPersistenceLayer() {
	url := util.Concat(
		mysqlConf.Username,
		":",
		mysqlConf.Password,
		"@tcp(",
		mysqlConf.Hostname,
		":",
		mysqlConf.Port,
		")/",
		mysqlConf.Database,
		"?charset=",
		mysqlConf.Charset,
		"&parseTime=",
		mysqlConf.ParseTime,
		"&loc=",
		mysqlConf.Loc,
	)
	var err error
	mysqlEngine, err = xorm.NewEngine(mysqlConf.Driver, url)
	if err != nil {
		util.LaunchLogger("xorm", false)
		util.FatalLogger(err, "xorm.NewEngine()")
		return
	}
	util.LaunchLogger("xorm", true)
}

// GetMySQLEngine to get a (fake) singleton *xorm.Engine instance.
// Why fake? Because InitPersistenceLayer will only be handled once in the runtime.
func GetMySQLEngine() *xorm.Engine {
	return mysqlEngine
}

// InitCachingLayer to connect redis database
func InitCachingLayer() {
	addr := util.Concat(redisConf.Hostname, ":", redisConf.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConf.Password,
		DB:       redisConf.Database,
	})
}

// GetRedisClient to get a (fake) singleton *redis.Client instance
func GetRedisClient() *redis.Client {
	return redisClient
}
