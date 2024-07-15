package product

import (
	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type FactoryConsumer interface {
	NewUsecaseUpsertSearch(req entity.EventBusUpsertESRequest) usecase.Usecase
}

type factoryConsumerImpl struct {
	db persistent.DB
	se search.SearchEngine
}

func NewFactoryConsumer(db persistent.DB, se search.SearchEngine) FactoryConsumer {
	return &factoryConsumerImpl{
		db: db,
		se: se,
	}
}
