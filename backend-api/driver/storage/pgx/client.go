package pgx

import (
	"context"
	"fmt"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientImpl struct {
	pool *pgxpool.Pool
}

func New(cfg config.DBConfig) *clientImpl {
	fmt.Println(cfg.ConnStr, "config conn str")
	poolCfg, err := pgxpool.ParseConfig(cfg.ConnStr)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolCfg)
	if err != nil {
		panic(err)
	}

	return &clientImpl{
		pool: pool,
	}
}

func (c *clientImpl) Dial(_ context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (c *clientImpl) GetPool() *pgxpool.Pool {
	panic("not implemented") // TODO: Implement
}

func (c *clientImpl) Close() {
	panic("not implemented") // TODO: Implement
}
