package mapper

import (
	"context"
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

var ctx = context.Background()

func getEngine() *xorm.Engine {
	return conf.GetMySQLEngine()
}

func getClient() *redis.Client {
	return conf.GetRedisClient()
}
