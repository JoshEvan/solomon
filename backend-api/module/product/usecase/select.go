package product

import (
	"context"

	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
)

type selectUsecase struct {
	db persistent.DB
}

func (f *factoryImpl) NewUsecaseSelect() usecase.Usecase {
	return &selectUsecase{
		db: f.db,
	}
}

func (u *selectUsecase) Do(ctx context.Context) (ret interface{}, err error) {
	return u.db.GetAll(ctx)
}
