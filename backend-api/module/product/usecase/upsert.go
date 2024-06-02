package product

import (
	"context"

	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
)

type upsertUsecase struct {
	req entity.UpsertRequest
	db  persistent.DB
}

func (u *upsertUsecase) Do(ctx context.Context) (interface{}, error) {

	return nil, nil
}
