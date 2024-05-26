package http

import (
	"context"
	"net/http"

	"github.com/JoshEvan/solomon/driver/net"
	"github.com/JoshEvan/solomon/driver/net/handler"
	product "github.com/JoshEvan/solomon/module/product/usecase"
)

type HTTPHandler interface {
	RegisterHandler(router net.Router)
}

type BaseHandler struct {
}

func InitHTTPHandler(router net.Router) {
	handlers := []HTTPHandler{
		CreateProductHandler(product.NewFactory()),
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
