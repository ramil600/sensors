package rabbit

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type AmqpDispatcher struct {
	queuename string
	channel   *amqp.Channel
}

func (ad AmqpDispatcher) Apply(ctx context.Context, command Command) error {
	//Build event from Command, command comes as payload from request
	evt := EventFromCommand(command)
	if evt == nil {
		log.Fatal("EventCommand: Event could not be parsed from command")
	}

	//Publish Event on the queue
	payload, err := json.Marshal(evt)
	if err != nil {
		log.Fatal("Apply: json.Marshal: Could not serialize Event Object")
	}

	err = ad.channel.Publish("",
		ad.queuename,
		false,
		false,
		amqp.Publishing{
			Body: payload,
		})
	if err != nil {
		log.Fatal("amqp.Channel.Publish could not publish event on the queue")
	}
	log.Println("Event is dispatched on ", ad.queuename)

	return nil
}

func NewDispatcher() Dispatcher {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &AmqpDispatcher{channel: channel}
}
