package test

import (
	"context"
	"fmt"
	"gin_gorm_oj/models"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func TestRedisSet(t *testing.T) {
	rdb.Set(ctx, "name", "12233", time.Second*5)
}
func TestRedistGet(t *testing.T) {
	v, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
func TestRedisGetByModels(t *testing.T) {
	v, err := models.RDB.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
