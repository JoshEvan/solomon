package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/JoshEvan/solomon/driver/storage"
	"github.com/JoshEvan/solomon/module/product/entity"
)

type Cache interface {
	Get(ctx context.Context, id string) (entity.Product, error)
	Upsert(context.Context, entity.Product) error
	Invalidate(context.Context, string) error
}

type cacheImpl struct {
	cache storage.Cache
}

func GetCache(cache storage.Cache) *cacheImpl {
	return &cacheImpl{
		cache: cache,
	}
}

const (
	cacheKeyProductTemplate = "product:%s"
	ttlMSecProduct          = time.Minute * 5
)

func (c *cacheImpl) Get(ctx context.Context, id string) (product entity.Product, err error) {
	ret, err := c.cache.Get(ctx, fmt.Sprintf(cacheKeyProductTemplate, id))
	if err != nil {
		log.Println(err.Error())
		return
	}
	if retVal, ok := ret.(string); ok {
		err = json.Unmarshal([]byte(retVal), &product)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	return
}

func (c *cacheImpl) Upsert(ctx context.Context, product entity.Product) error {
	val, err := json.Marshal(product)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = c.cache.Set(ctx, fmt.Sprintf(cacheKeyProductTemplate, product.Id), val, int(ttlMSecProduct.Milliseconds()))
	if err != nil {
		log.Println("UPSERT", err)
	}
	return err
}

func (c *cacheImpl) Invalidate(ctx context.Context, id string) error {
	return c.cache.Del(ctx, fmt.Sprintf(cacheKeyProductTemplate, id))
}
