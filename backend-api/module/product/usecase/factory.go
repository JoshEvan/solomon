package product

import (
	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/cache"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type Factory interface {
	NewUsecaseUpsert(req entity.UpsertRequest) usecase.Usecase
	NewUsecaseSelect(req entity.SelectQuery) usecase.Usecase
}

type factoryImpl struct {
	db       persistent.DB
	se       search.SearchEngine
	cache    cache.Cache
	eventBus bus.EventBusProducer
}

func NewFactory(db persistent.DB, se search.SearchEngine, cache cache.Cache, eventBus bus.EventBusProducer) Factory {
	return &factoryImpl{
		db:       db,
		se:       se,
		cache:    cache,
		eventBus: eventBus,
	}
}
