package rabbit

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type Dispatcher interface {
	Apply(context.Context, Command) error
}

type Publisher interface {
	Publish(string, string, bool, bool, amqp.Publishing) error
}

type AmqpDispatcher struct {
	Queuename string
	Conn      *amqp.Connection
	Channel   *amqp.Channel
}

func (ad AmqpDispatcher) Apply(ctx context.Context, command Command) error {
	//Build event from Command, command comes as payload from request
	evt := EventFromCommand(command)
	if evt == nil {
		log.Fatal("EventCommand: Event could not be parsed from command")
	}

	q, err := ad.Channel.QueueDeclare(
		"",    //queue name
		false, //durable
		false, //AutoDelete
		false, //exclusive
		false, //noWait
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	ad.Queuename = q.Name

	//Publish Event on the queue
	payload, err := json.Marshal(evt)
	if err != nil {
		log.Fatal("Apply: json.Marshal: Could not serialize Event Object")
	}

	err = ad.Channel.Publish("",
		ad.Queuename,
		false,
		false,
		amqp.Publishing{
			Body: payload,
		})
	if err != nil {
		log.Fatal("amqp.Channel.Publish could not publish event on the queue")
	}
	log.Println("Event is dispatched on ", ad.Queuename)

	return nil
}

func NewDispatcher(url string) *AmqpDispatcher {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	return &AmqpDispatcher{Channel: channel,
		Conn: conn}
}
