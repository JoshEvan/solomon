package storage

import (
	"context"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/storage/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Dial(context.Context) error
	GetPool() *pgxpool.Pool
	Close()
}

type DB interface {
	Execute(ctx context.Context, query string, params ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, params ...interface{}) error
}

func New(cfg config.DBConfig) Client {
	switch cfg.Driver {
	case config.Pgx:
		return pgx.New(cfg)
	}
	return nil
}
