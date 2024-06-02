package pgx

import "context"

type pgxImpl struct {
	client *clientImpl
}

func NewDB(cli *clientImpl) *pgxImpl {
	return &pgxImpl{
		client: cli,
	}
}

func (d *pgxImpl) Execute(ctx context.Context, query string, params ...interface{}) error {
	panic("not implemented") // TODO: Implement
}

func (d *pgxImpl) Select(ctx context.Context, dest interface{}, query string, params ...interface{}) error {
	panic("not implemented") // TODO: Implement
}
