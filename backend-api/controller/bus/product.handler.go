package bus

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/JoshEvan/solomon/driver/bus"
	"github.com/JoshEvan/solomon/module/product/entity"
	product "github.com/JoshEvan/solomon/module/product/usecase"
	gonsq "github.com/nsqio/go-nsq"
)

type ProductHandler struct {
	Usecase product.FactoryConsumer
}

func CreateProductHandler(u product.FactoryConsumer) ProductHandler {
	return ProductHandler{
		Usecase: u,
	}
}

func (p *ProductHandler) RegisterConsumerHandlers(consumers *bus.EventBusConsumer) {
	(*consumers).RegisterHandler(entity.EventUpsertES, func(message interface{}) error {
		ctx := context.Background()
		messageNSQ, ok := message.(*gonsq.Message)
		if !ok {
			err := errors.New("invalid message")
			log.Println(err.Error())
			return err
		}
		req := entity.EventBusUpsertESRequest{}
		err := json.Unmarshal(messageNSQ.Body, &req)
		if err != nil {
			log.Println(err.Error())
			messageNSQ.Finish()
			return err
		}
		_, err = p.Usecase.NewUsecaseUpsertSearch(req).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
