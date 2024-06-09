package persistent

import (
	"context"
	"log"

	"github.com/JoshEvan/solomon/driver/storage"
	"github.com/JoshEvan/solomon/module/product/entity"
	"github.com/google/uuid"
)

type DB interface {
	Insert(context.Context, entity.Product) (uuid.UUID, error)
	Update(context.Context, entity.Product) error
	GetAll(context.Context) ([]entity.Product, error)
	Get(context.Context, uuid.UUID) (entity.Product, error)
}

type dbImpl struct {
	db storage.DB
}

func GetDB(db storage.DB) *dbImpl {
	return &dbImpl{
		db: db,
	}
}

func (p *dbImpl) Insert(ctx context.Context, data entity.Product) (id uuid.UUID, err error) {
	err = p.db.ExecuteAndScan(ctx, &id, insertQueryProduct, data.Name, data.ImgUrl)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *dbImpl) Update(ctx context.Context, data entity.Product) (err error) {
	err = p.db.Execute(ctx, updateQueryProduct, data.Name, data.ImgUrl, data.Id)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *dbImpl) GetAll(ctx context.Context) (result []entity.Product, err error) {
	err = p.db.Select(ctx, &result, selectAllQueryProduct)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

func (p *dbImpl) Get(ctx context.Context, id uuid.UUID) (result entity.Product, err error) {
	err = p.db.Get(ctx, &result, selectByIdQueryProduct, id)
	if err != nil {
		log.Println(err.Error())
	}
	return
}
