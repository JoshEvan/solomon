package product

import (
	"context"
	"errors"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/cache"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type selectUsecase struct {
	db    persistent.DB
	se    search.SearchEngine
	cache cache.Cache
	req   entity.SelectQuery
}

func (f *factoryImpl) NewUsecaseSelect(req entity.SelectQuery) usecase.Usecase {
	return &selectUsecase{
		db:    f.db,
		se:    f.se,
		cache: f.cache,
		req:   req,
	}
}

func (u *selectUsecase) Validate(ctx context.Context) error {
	if u.req.Page < 0 || u.req.Limit < 0 {
		return errors.New("invalid request")
	}
	return nil
}

func (u *selectUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	if err := u.Validate(ctx); err != nil {
		return nil, err
	}

	found, err := u.se.Search(ctx, entity.SearchProductRequest{
		SearchText: u.req.SearchText,
		PriceMin:   u.req.PriceMin,
		PriceMax:   u.req.PriceMax,
	}, u.req.Page, u.req.Limit, u.req.SortBy, u.req.IsSortAsc)
	if err != nil {
		return nil, err
	}
	if ids := found.GetIds(); len(ids) > 0 {
		data := make([]entity.Product, len(ids))
		notCachedIds := []string{}
		notCachedIdxs := []int{}
		for i, id := range ids {
			cached, err := u.cache.Get(ctx, id)
			if err == nil {
				data[i] = cached
			} else {
				notCachedIds = append(notCachedIds, id)
				notCachedIdxs = append(notCachedIdxs, i)
			}
		}

		if len(notCachedIds) > 0 {
			dbData, err := u.db.GetBulkIds(ctx, ids)
			if err != nil {
				return nil, err
			}
			for i, product := range dbData {
				u.cache.Upsert(ctx, product)
				data[notCachedIdxs[i]] = product
			}
		}
		return data, nil
	}
	return nil, nil
}
