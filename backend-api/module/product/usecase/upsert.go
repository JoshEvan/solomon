package product

import (
	"context"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/google/uuid"
)

type upsertUsecase struct {
	req entity.UpsertRequest
	db  persistent.DB
}

func (f *factoryImpl) NewUsecaseUpsert(req entity.UpsertRequest) usecase.Usecase {
	return &upsertUsecase{
		db:  f.db,
		req: req,
	}
}

func (u *upsertUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	var existed entity.Product
	if u.req.Id != uuid.Nil {
		existed, err = u.db.Get(ctx, u.req.Id)
		if err != nil {
			return nil, err
		}
	}
	if existed.Id == uuid.Nil {
		id, err := u.db.Insert(ctx, entity.Product(u.req))
		if err != nil {
			return nil, err
		}
		u.req.Id = id
	} else {
		err := u.db.Update(ctx, entity.Product(u.req))
		if err != nil {
			return nil, err
		}
	}
	return u.req, nil
}
