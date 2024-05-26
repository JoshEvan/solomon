package http

import (
	"context"
	"net/http"

	"github.com/JoshEvan/solomon/driver/net"
	product "github.com/JoshEvan/solomon/module/product/usecase"
)

type ProductHandler struct {
	BaseHandler
	Usecase product.Factory
}

func CreateProductHandler(u product.Factory) HTTPHandler {
	return &ProductHandler{
		BaseHandler: BaseHandler{},
		Usecase:     u,
	}
}

func (p *ProductHandler) RegisterHandler(router net.Router) {
	subrouter := router.RegisterSubRouter("/product")
	subrouter.RegisterHandler("/upsert", p.BaseHandler.Handle(p.upsert), http.MethodPost)
	// subrouter.RegisterHandler("/list", p.Usecase.NewShowHandler, http.MethodGet)
}

func (p *ProductHandler) upsert(ctx context.Context, r *http.Request) (interface{}, error) {
	p.Usecase.NewUsecaseUpsert().Do(context.Background())
	return nil, nil
}
