package main

import (
	"common"
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
	listen(ch)
}

func listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			o := &common.Order{}
			if err := json.Unmarshal(d.Body, o); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", err)
				continue
			}
			paymentLink, err := createPaymentLink()
			if err != nil {
				log.Printf("failed to create payment: %v", err)
				// handle retry here ...
				continue
			}
			log.Printf("Payment link generated: %s", paymentLink)
		}
	}()
	log.Printf("AMQP Listening. To exit press CTRL+C")
	<-forever
}

func createPaymentLink() (string, error) {
	return "dummy-payment-link.com", nil
}
