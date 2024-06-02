package product

import (
	"github.com/JoshEvan/solomon/driver/usecase"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
)

type Factory interface {
	NewUsecaseUpsert() usecase.Usecase
	// NewShowHandler(ctx context.Context, r *http.Request) (interface{}, error)
}

type factoryImpl struct {
	db persistent.DB
}

func NewFactory(db persistent.DB) Factory {
	return &factoryImpl{
		db: db,
	}
}

func (f *factoryImpl) NewUsecaseUpsert() usecase.Usecase {
	return &upsertUsecase{}
}
