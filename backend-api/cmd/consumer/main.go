package main

import (
	"os"
	"os/signal"
	"syscall"

	controller "github.com/JoshEvan/solomon/controller/bus"
	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/storage/elastic"
	"github.com/JoshEvan/solomon/driver/storage/entity"
	"github.com/JoshEvan/solomon/driver/storage/pgx"
	"github.com/JoshEvan/solomon/driver/util"
	"github.com/JoshEvan/solomon/module/product/repository/persistent"
	"github.com/JoshEvan/solomon/module/product/repository/search"
	product "github.com/JoshEvan/solomon/module/product/usecase"
)

func main() {
	defer util.PanicCapture()

	cfg := config.Get()
	consumers := bus.NewEventConsumers(config.Get().Consumer)
	handler := controller.CreateProductHandler(product.NewFactoryConsumer(
		persistent.GetDB(pgx.NewDB(pgx.New(cfg.DBConfig))),
		search.GetSearchEngine(elastic.New(entity.Config(cfg.SearchConfig))),
	))
	handler.RegisterConsumerHandlers(&consumers)
	consumers.Connect()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumers.Stop()
}
