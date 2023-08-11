package mapper

import (
	"context"
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/redis/go-redis/v9"
	"time"
	"xorm.io/xorm"
)

var defaultTimeout = 5

func getContext(timeout int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	return ctx, cancel
}

func getDefaultContext() (context.Context, context.CancelFunc) {
	return getContext(defaultTimeout)
}

func getEngine() *xorm.Engine {
	return conf.GetMySQLEngine()
}

func getClient() *redis.Client {
	return conf.GetRedisClient()
}
