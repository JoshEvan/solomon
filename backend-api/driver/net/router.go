package net

import (
	"net/http"

	"github.com/JoshEvan/solomon/driver/net/handler"
	"github.com/JoshEvan/solomon/driver/net/router/mux"
)

type Router interface {
	GetHTTPHandler() http.Handler
	RegisterSubRouter(pathPrefix string) handler.SubRouter
}

func NewRouter() Router {
	return mux.NewRouter()
}
