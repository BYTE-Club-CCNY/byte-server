package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)
var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
}

func GetCache(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()

	if err != nil {
		return val, nil
	}
	return val, err
}

func DeleteFromCache(key string) error {
	res := rdb.Del(ctx, key)
	return res.Err()
}

func AddToCache(key string, value string, timeout time.Duration) error {
	err := rdb.Set(ctx, key, value, timeout).Err()
	return err
}

func ClearCache() error {
	err := rdb.FlushAll(ctx)

	if err.Err() != nil {
		return err.Err()
	}

	return nil
}