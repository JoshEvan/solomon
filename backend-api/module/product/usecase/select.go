package product

import (
	"context"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type selectUsecase struct {
	db  persistent.DB
	se  search.SearchEngine
	req entity.SelectQuery
}

func (f *factoryImpl) NewUsecaseSelect(req entity.SelectQuery) usecase.Usecase {
	return &selectUsecase{
		db:  f.db,
		se:  f.se,
		req: req,
	}
}

func (u *selectUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	found, err := u.se.Search(ctx, entity.SearchProductRequest{
		SearchText: u.req.SearchText,
		PriceMin:   u.req.PriceMin,
		PriceMax:   u.req.PriceMax,
	}, u.req.Page, u.req.Limit, u.req.SortBy, u.req.IsSortAsc)
	if err != nil {
		return nil, err
	}
	if ids := found.GetIds(); len(ids) > 0 {
		return u.db.GetBulkIds(ctx, ids)
	}
	return nil, nil
}
