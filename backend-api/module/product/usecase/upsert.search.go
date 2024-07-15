package product

import (
	"context"
	"log"

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

func (u *upsertUsecaseSearch) Validate(ctx context.Context) (err error) {
	return nil
}

func (u *upsertUsecaseSearch) Do(ctx context.Context) (ret interface{}, err error) {
	existed, err := u.db.GetBulkIds(ctx, []string{u.req.Id})
	if err != nil {
		return nil, err
	}
	if len(existed) == 0 {
		return nil, nil
	}
	dbValue := existed[0]
	if dbValue.UpdateTime.After(u.req.Timestamp) {
		log.Println("outdated search value")
		return nil, nil
	}

	if !u.req.IsUpdate {
		err = u.se.Insert(ctx, entity.IndexedProduct{
			Id:    u.req.Id,
			Name:  u.req.Name,
			Price: u.req.Price,
		})
		if err != nil {
			return nil, err
		}
	} else {
		err = u.se.Update(ctx, entity.IndexedProduct{
			Id:    u.req.Id,
			Name:  u.req.Name,
			Price: u.req.Price,
		})
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
