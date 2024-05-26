package mux

import (
	"net/http"

	"github.com/JoshEvan/solomon/driver/net/handler"
	"github.com/gorilla/mux"
)

type RouterMux struct {
	router     *mux.Router
	subrouters []handler.SubRouter
}

func NewRouter() *RouterMux {
	return &RouterMux{
		router: mux.NewRouter(),
	}
}

func (r *RouterMux) GetHTTPHandler() http.Handler {
	return r.router
}

// func (r *RouterMux) RegisterSubRouter(pathPrefix string) net.SubRouter {
// 	newSubrouter := handler.NewSubRouter(r.router.PathPrefix(pathPrefix).Subrouter())
// 	r.subRoutes = append(r.subRoutes, newSubrouter)
// 	return newSubrouter
// }

func (r *RouterMux) RegisterSubRouter(pathPrefix string) handler.SubRouter {
	newSubrouter := handler.NewHTTPHandler(
		pathPrefix, r.router.PathPrefix(pathPrefix).Subrouter(),
	)
	r.subrouters = append(r.subrouters, newSubrouter)
	return newSubrouter
}
