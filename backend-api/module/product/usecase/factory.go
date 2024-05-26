package product

import (
	"github.com/JoshEvan/solomon/driver/usecase"
)

type Factory interface {
	NewUsecaseUpsert() usecase.Usecase
	// NewShowHandler(ctx context.Context, r *http.Request) (interface{}, error)
}

type factoryImpl struct {
	// db
}

func NewFactory() Factory {
	return &factoryImpl{}
}

func (f *factoryImpl) NewUsecaseUpsert() usecase.Usecase {
	return &upsertUsecase{}
}
