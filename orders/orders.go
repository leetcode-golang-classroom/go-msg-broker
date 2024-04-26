package main

import (
	"common"
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	if common.C == nil || common.C.RABBITMQ_URL == "" {
		log.Fatal("failed to config")
	}
	ch, close := common.ConnectAmqp(common.C.RABBITMQ_URL)
	defer func() {
		close()
		ch.Close()
	}()

	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	marshalledOrder, err := json.Marshal(common.Order{
		ID: "order-2",
		Items: []common.Item{
			{
				ID:       "item-2",
				Quantity: 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	err = ch.PublishWithContext(context.Background(), "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        marshalledOrder,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Order published")
}
