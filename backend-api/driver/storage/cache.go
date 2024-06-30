package storage

import (
	"context"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/storage/redis"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, val interface{}, ttlMSec int) error
}

func NewCache(cfg config.CacheConfig) Cache {
	return redis.New(cfg)
}
