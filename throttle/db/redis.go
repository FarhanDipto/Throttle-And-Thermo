package db

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379", 
	})
}

func CheckRate(apiKey, action string, limit, windowSeconds int) (int, bool) {
	key := fmt.Sprintf("rate:%s:%s", apiKey, action)

	cnt, _ := Rdb.Incr(ctx, key).Result()

	if cnt == 1 {
		Rdb.Expire(ctx, key, time.Duration(windowSeconds)*time.Second)
	}

	allowed := cnt <= int64(limit)
	remaining := limit - int(cnt)

	if remaining < 0 {
		remaining = 0
	}
	return remaining, allowed
}
