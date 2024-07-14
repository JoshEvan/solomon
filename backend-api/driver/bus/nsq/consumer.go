package nsq

import (
	"errors"
	"log"

	"github.com/JoshEvan/solomon/driver/config"
	"github.com/nsqio/go-nsq"
	gonsq "github.com/nsqio/go-nsq"
)

type consumers struct {
	nsqlookupAddress string
	client           map[config.EventBusTopic]*gonsq.Consumer
}

func NewConsumers(cfg config.EventBusConsumerConfig) *consumers {
	consumersNSQ := map[config.EventBusTopic]*gonsq.Consumer{}
	configNSQ := nsq.NewConfig()
	for _, v := range cfg.Listen {
		consumerNSQ, err := gonsq.NewConsumer(string(v.Topic), v.Channel, configNSQ)
		if err != nil {
			panic(err)
		}
		consumersNSQ[v.Topic] = consumerNSQ
	}
	return &consumers{
		nsqlookupAddress: cfg.ListenAddress,
		client:           consumersNSQ,
	}
}

func (c *consumers) Connect() {
	for _, v := range c.client {
		err := v.ConnectToNSQLookupd(c.nsqlookupAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *consumers) Stop() {
	for _, v := range c.client {
		v.Stop()
	}
}

func (c *consumers) RegisterHandler(topic config.EventBusTopic, f func(message interface{}) error) error {
	if _, ok := c.client[topic]; ok {
		c.client[topic].AddHandler(gonsq.HandlerFunc(func(m *gonsq.Message) error {
			return f(m)
		}))
	}
	return errors.New("topic not found")
}
