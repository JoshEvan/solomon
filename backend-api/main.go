package main

import (
	"github.com/JoshEvan/solomon/controller/http"
	"github.com/JoshEvan/solomon/driver/net"
)

func main() {
	httpRouter := net.NewRouter()
	http.InitHTTPHandler(httpRouter)
	err := net.ServeHTTP(httpRouter.GetHTTPHandler(), ":9099")
	if err != nil {
		panic(err)
	}
}