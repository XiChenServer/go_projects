package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
	"virus/global"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}
func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       0,
		PoolSize: redisConf.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败 %s", redisConf.Addr())
		return nil
	}
	return rdb
}
