package pgx

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
)

type pgxImpl struct {
	client *clientImpl
}

func NewDB(cli *clientImpl) *pgxImpl {
	return &pgxImpl{
		client: cli,
	}
}

func (d *pgxImpl) Execute(ctx context.Context, query string, params ...interface{}) error {
	_, err := d.client.GetPool().Exec(ctx, query, params...)
	return err
}

func (d *pgxImpl) ExecuteAndScan(ctx context.Context, dest interface{}, query string, params ...interface{}) error {
	return pgxscan.Get(ctx, d.client.GetPool(), dest, query, params...)
}

func (d *pgxImpl) Select(ctx context.Context, dest interface{}, query string, params ...interface{}) error {
	return pgxscan.Select(ctx, d.client.GetPool(), dest, query, params...)
}

func (d *pgxImpl) Get(ctx context.Context, dest interface{}, query string, params ...interface{}) error {
	return pgxscan.Get(ctx, d.client.GetPool(), dest, query, params...)
}
