package redis

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb  *redis.Client
	once sync.Once
	ctx  = context.Background()
)

func getClient() *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		})
	})
	return rdb
}
