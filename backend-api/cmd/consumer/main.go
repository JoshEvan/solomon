package main

import (
	"os"
	"os/signal"
	"syscall"

	controller "github.com/JoshEvan/solomon/controller/bus"
	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/driver/config"
	"github.com/JoshEvan/solomon/driver/util"
)

func main() {
	defer util.PanicCapture()

	consumers := bus.NewEventConsumers(config.Get().Consumer)
	controller.RegisterConsumerHandlers(&consumers)
	consumers.Connect()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	consumers.Stop()
}
