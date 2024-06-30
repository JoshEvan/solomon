package redis

import (
	"context"
	"time"

	"github.com/JoshEvan/solomon/driver/config"
	goredis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *goredis.Client
}

func New(cfg config.CacheConfig) *RedisClient {
	client := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
	})
	return &RedisClient{
		client,
	}
}

func (r *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Set(ctx context.Context, key string, val interface{}, ttlMSec int) error {
	return r.client.Set(ctx, key, val, time.Duration(ttlMSec)*time.Millisecond).Err()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
