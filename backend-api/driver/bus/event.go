package bus

import (
	"context"

	"github.com/JoshEvan/solomon/driver/bus/nsq"
	"github.com/JoshEvan/solomon/driver/config"
)

type EventBusProducer interface {
	Produce(ctx context.Context, event string, message interface{}) error
}

func NewEventPublisher(cfg config.EventBusConfig) EventBusProducer {
	return nsq.New(cfg)
}
