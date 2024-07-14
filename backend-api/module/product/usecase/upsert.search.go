package product

import (
	"context"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type upsertUsecaseSearch struct {
	req entity.EventBusUpsertESRequest
	db  persistent.DB
	se  search.SearchEngine
}

func (f *factoryConsumerImpl) NewUsecaseUpsertSearch(req entity.EventBusUpsertESRequest) usecase.Usecase {
	return &upsertUsecaseSearch{
		db:  f.db,
		se:  f.se,
		req: req,
	}
}

func (u *upsertUsecaseSearch) Do(ctx context.Context) (ret interface{}, err error) {
	panic("not implemented")
}
