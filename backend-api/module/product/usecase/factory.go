package product

import (
	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
)

type Factory interface {
	NewUsecaseUpsert(req entity.UpsertRequest) usecase.Usecase
	NewUsecaseSelect() usecase.Usecase
}

type factoryImpl struct {
	db persistent.DB
}

func NewFactory(db persistent.DB) Factory {
	return &factoryImpl{
		db: db,
	}
}
