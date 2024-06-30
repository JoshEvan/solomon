package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoshEvan/solomon/driver/net"
	"github.com/JoshEvan/solomon/driver/util"
	"github.com/JoshEvan/solomon/module/product/entity"
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
	subrouter.RegisterHandler("/list", p.BaseHandler.Handle(p.index), http.MethodGet)
}

func (p *ProductHandler) upsert(ctx context.Context, r *http.Request) (interface{}, error) {
	req := entity.UpsertRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("error parsing request", err.Error())
		return nil, err
	}
	return p.Usecase.NewUsecaseUpsert(req).Do(ctx)
}

func (p *ProductHandler) index(ctx context.Context, r *http.Request) (interface{}, error) {
	req := entity.SelectQuery{
		Page:       util.StringToIntUnsafe(r.FormValue(entity.PageFormValue)),
		Limit:      util.StringToIntUnsafe(r.FormValue(entity.LimitFormValue)),
		SearchText: r.FormValue(entity.SearchTextFormValue),
		SortBy:     r.FormValue(entity.SortByFormValue),
		IsSortAsc:  util.StringToIntUnsafe(r.FormValue(entity.IsSortAscFormValue)) == 1,
		PriceMin:   util.StringToFloat64Unsafe(r.FormValue(entity.PriceMinFormValue)),
		PriceMax:   util.StringToFloat64Unsafe(r.FormValue(entity.PriceMaxFormValue)),
	}
	return p.Usecase.NewUsecaseSelect(req).Do(ctx)
}
