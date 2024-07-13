package nsq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/JoshEvan/solomon/driver/config"
	gonsq "github.com/nsqio/go-nsq"
)

type nsqPublisher struct {
	producer *gonsq.Producer
}

func New(cfg config.EventBusConfig) *nsqPublisher {
	producer, err := gonsq.NewProducer(cfg.PublishAddress, gonsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}
	return &nsqPublisher{producer}
}

func (n *nsqPublisher) Produce(ctx context.Context, event string, message interface{}) error {
	messageBody, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return n.producer.Publish(event, messageBody)
}
