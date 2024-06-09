package pgx

import (
	"context"
	"log"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientImpl struct {
	pool *pgxpool.Pool
}

func New(cfg config.DBConfig) *clientImpl {
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

func (c *clientImpl) Dial(ctx context.Context) error {
	if err := c.pool.Ping(ctx); err != nil {
		log.Println("[DB] ping error", err.Error())
		return err
	}
	return nil
}

func (c *clientImpl) GetPool() *pgxpool.Pool {
	return c.pool
}

func (c *clientImpl) Close() {
	c.pool.Close()
}
