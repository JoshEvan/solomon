package http

import (
	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/net"
	"github.com/JoshEvan/solomon/driver/storage"
	"github.com/JoshEvan/solomon/driver/storage/elastic"
	"github.com/JoshEvan/solomon/driver/storage/entity"
	"github.com/JoshEvan/solomon/driver/storage/pgx"
	"github.com/JoshEvan/solomon/module/product/repository/cache"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
	product "github.com/JoshEvan/solomon/module/product/usecase"
)

type HTTPHandler interface {
	RegisterHandler(router net.Router)
}

type BaseHandler struct {
}

func InitHTTPHandler(router net.Router, cfg config.Config) {
	handlers := []HTTPHandler{
		CreateProductHandler(
			product.NewFactory(
				persistent.GetDB(pgx.NewDB(pgx.New(cfg.DBConfig))),
				search.GetSearchEngine(elastic.New(entity.Config(cfg.SearchConfig))),
				cache.GetCache(storage.NewCache(cfg.CacheConfig)),
				bus.NewEventPublisher(cfg.EventBusConfig),
			),
		),
	}

	for _, handler := range handlers {
		handler.RegisterHandler(router)
	}
}
