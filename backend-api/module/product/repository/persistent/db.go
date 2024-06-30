package persistent

import (
	"context"
	"log"

	"github.com/JoshEvan/solomon/driver/storage"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/lib/pq"
)

type DB interface {
	Insert(context.Context, entity.Product) (string, error)
	Update(context.Context, entity.Product) error
	GetAll(context.Context) ([]entity.Product, error)
	GetBulkIds(context.Context, []string) ([]entity.Product, error)
}

type dbImpl struct {
	db storage.DB
}

func GetDB(db storage.DB) *dbImpl {
	return &dbImpl{
		db: db,
	}
}

func (p *dbImpl) Insert(ctx context.Context, data entity.Product) (id string, err error) {
	err = p.db.ExecuteAndScan(ctx, &id, insertQueryProduct, data.Name, data.ImgUrl, data.Price)
	if err != nil {
		log.Println(insertQueryProduct, err.Error())
	}
	return
}

func (p *dbImpl) Update(ctx context.Context, data entity.Product) (err error) {
	err = p.db.Execute(ctx, updateQueryProduct, data.Name, data.ImgUrl, data.Price, data.Id)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *dbImpl) GetAll(ctx context.Context) (result []entity.Product, err error) {
	err = p.db.Select(ctx, &result, selectAllQueryProduct)
	if err != nil {
		log.Println(selectAllQueryProduct, err.Error())
	}
	return
}

func (p *dbImpl) GetBulkIds(ctx context.Context, ids []string) (result []entity.Product, err error) {
	err = p.db.Select(ctx, &result, selectByIdQueryProduct, pq.Array(ids))
	if err != nil {
		log.Println(selectByIdQueryProduct, err.Error())
	}
	return
}
