package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go_ranking/config"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	Rctx = context.Background()
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
