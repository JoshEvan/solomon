package product

import (
	"context"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/cache"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type upsertUsecase struct {
	req   entity.UpsertRequest
	db    persistent.DB
	cache cache.Cache
	se    search.SearchEngine
}

func (f *factoryImpl) NewUsecaseUpsert(req entity.UpsertRequest) usecase.Usecase {
	return &upsertUsecase{
		db:    f.db,
		se:    f.se,
		cache: f.cache,
		req:   req,
	}
}

func (u *upsertUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	var existed []entity.Product
	if u.req.Id != "" {
		existed, err = u.db.GetBulkIds(ctx, []string{u.req.Id})
		if err != nil {
			return nil, err
		}
	}
	if len(existed) == 0 {
		id, err := u.db.Insert(ctx, entity.Product(u.req))
		if err != nil {
			return nil, err
		}
		u.req.Id = id
		err = u.se.Insert(ctx, entity.IndexedProduct{
			Id:    u.req.Id,
			Name:  u.req.Name,
			Price: u.req.Price,
		})
		if err != nil {
			return nil, err
		}
	} else {
		err := u.db.Update(ctx, entity.Product(u.req))
		if err != nil {
			return nil, err
		}
		if err := u.cache.Invalidate(ctx, u.req.Id); err != nil {
			return nil, err
		}
		err = u.se.Update(ctx, entity.IndexedProduct{
			Id:    u.req.Id,
			Name:  u.req.Name,
			Price: u.req.Price,
		})
		if err != nil {
			return nil, err
		}
	}
	return u.req, nil
}
