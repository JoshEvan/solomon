package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/JoshEvan/solomon/driver/net/handler"
	"github.com/JoshEvan/solomon/driver/util"
)

func (b *BaseHandler) Handle(handlerFunc handler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer util.PanicCapture()
		w.Header().Set("Content-Type", "application/json")
		response, err := handlerFunc(context.Background(), r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		jsonResponse, _ := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
