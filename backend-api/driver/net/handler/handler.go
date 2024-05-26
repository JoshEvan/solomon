package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SubRouter interface {
	RegisterHandler(path string, f func(http.ResponseWriter, *http.Request), methods ...string)
}

type HTTPHandlerMux struct {
	r          *mux.Router
	pathPrefix string
}

type HandlerFunc func(context.Context, *http.Request) (interface{}, error)

func NewHTTPHandler(pathPrefix string, subrouter *mux.Router) *HTTPHandlerMux {
	return &HTTPHandlerMux{
		pathPrefix: pathPrefix,
		r:          subrouter,
	}
}

func (h *HTTPHandlerMux) RegisterHandler(path string, f func(http.ResponseWriter, *http.Request), methods ...string) {
	log.Printf("Register listener on %s%s", h.pathPrefix, path)
	h.r.HandleFunc(path, f).Methods(methods...)
}
