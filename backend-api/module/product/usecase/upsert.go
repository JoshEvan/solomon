package product

import (
	"context"
	"errors"
	"time"

	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/cache"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
)

type upsertUsecase struct {
	req      entity.UpsertRequest
	db       persistent.DB
	cache    cache.Cache
	se       search.SearchEngine
	eventBus bus.EventBusProducer
}

func (f *factoryImpl) NewUsecaseUpsert(req entity.UpsertRequest) usecase.Usecase {
	return &upsertUsecase{
		db:       f.db,
		se:       f.se,
		cache:    f.cache,
		eventBus: f.eventBus,
		req:      req,
	}
}

func (u *upsertUsecase) Validate(ctx context.Context) (err error) {
	if u.req.Name == "" || u.req.Price < 0 {
		return errors.New("invalid request")
	}
	return nil
}

func (u *upsertUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	if err := u.Validate(ctx); err != nil {
		return nil, err
	}

	var existed []entity.Product
	if u.req.Id != "" {
		existed, err = u.db.GetBulkIds(ctx, []string{u.req.Id})
		if err != nil {
			return nil, err
		}
	}
	isUpdate := len(existed) == 0
	if !isUpdate {
		id, err := u.db.Insert(ctx, u.req.ToProduct())
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
		err := u.db.Update(ctx, u.req.ToProduct())
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
	u.eventBus.Produce(ctx, entity.EventUpsertES, entity.EventBusUpsertESRequest{
		UpsertRequest: u.req,
		Timestamp:     time.Now(),
		IsUpdate:      isUpdate,
	})
	return u.req, nil
}
