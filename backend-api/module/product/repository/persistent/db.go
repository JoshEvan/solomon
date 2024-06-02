package persistent

import (
	"context"
	"log"

	"github.com/JoshEvan/solomon/driver/storage"
	"github.com/JoshEvan/solomon/module/product/entity"
)

type DB interface {
	Upsert(context.Context, entity.Product) error
	GetAll(context.Context) ([]entity.Product, error)
}

type dbImpl struct {
	db storage.DB
}

func GetDB(db storage.DB) *dbImpl {
	return &dbImpl{
		db: db,
	}
}

func (p *dbImpl) Upsert(ctx context.Context, data entity.Product) (err error) {
	err = p.db.Execute(ctx, upsertQueryProduct,
		data.Id, data.Name, data.ImgUrl)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *dbImpl) GetAll(ctx context.Context) (result []entity.Product, err error) {
	err = p.db.Select(ctx, result, selectAllQueryProduct)
	if err != nil {
		log.Println(err.Error())
	}
	return
}
