package http

import (
	"context"
	"net/http"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/net"
	"github.com/JoshEvan/solomon/driver/net/handler"
	"github.com/JoshEvan/solomon/driver/storage/pgx"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	product "github.com/JoshEvan/solomon/module/product/usecase"
)

type HTTPHandler interface {
	RegisterHandler(router net.Router)
}

type BaseHandler struct {
}

func InitHTTPHandler(router net.Router, cfg config.Config) {
	handlers := []HTTPHandler{
		CreateProductHandler(
			product.NewFactory(persistent.GetDB(
				pgx.NewDB(pgx.New(cfg.DBConfig)),
			)),
		),
	}

	for _, handler := range handlers {
		handler.RegisterHandler(router)
	}
}

func (b *BaseHandler) Handle(handlerFunc handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(context.Background(), r)
	}
}