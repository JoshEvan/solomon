package bus

import (
	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/module/product/entity"
)

func RegisterConsumerHandlers(consumers *bus.EventBusConsumer) {
	(*consumers).RegisterHandler(entity.EventUpsertES, func(message interface{}) error {
		return nil
	})
}
