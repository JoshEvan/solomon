package bus

import (
	"context"

	"github.com/JoshEvan/solomon/driver/bus/nsq"
	"github.com/JoshEvan/solomon/driver/config"
)

type EventBusProducer interface {
	Produce(ctx context.Context, event string, message interface{}) error
}

type EventBusConsumer interface {
	Connect()
	Stop()
	RegisterHandler(topic config.EventBusTopic, f func(message interface{}) error) error
}

// type Handler interface {
// 	HandleMessage(message interface{}) error
// }

func NewEventPublisher(cfg config.EventBusConfig) EventBusProducer {
	return nsq.New(cfg)
}

func NewEventConsumers(cfg config.EventBusConsumerConfig) EventBusConsumer {
	return nsq.NewConsumers(cfg)
}
