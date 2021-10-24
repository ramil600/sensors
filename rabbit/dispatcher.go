package rabbit

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"time"
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

// Apply receives command that it needs to update Sensor(C_UD) and fires event on rabbit queue
// later it can be refactored to receive different objects for this purpose
// to implement: fires events after the changes to repo applied
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

// NewDispatcher returns new dispatcher that subscribes to queue of commands, *todo*(events)
func NewDispatcher(url string) *AmqpDispatcher {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	q, err := channel.QueueDeclare(
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
	ad := &AmqpDispatcher{
		Queuename: q.Name,
		Channel:   channel,
		Conn:      conn}

	log.Println("Dispatcher created with Queue Name:", ad.Queuename)

	return ad
}

// EventFromCommand is helper to produce an event from command
func EventFromCommand(command Command) Event {

	switch v := command.(type) {
	case CreateSensor:
		cmd := command.(CreateSensor)
		evt := SensorCreated{
			Name:       cmd.Name,
			Sensortype: cmd.Sensortype,
			EventModel: EventModel{
				ID:        cmd.GetId(),
				Version:   0,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		log.Println("Command type captured is: ", v)
		return evt
	case UpdateSensor:
		cmd := command.(UpdateSensor)
		evt := SensorUpdated{
			Name: cmd.Name,
			Sensortype: cmd.Sensortype,
			EventModel: EventModel{
				ID: cmd.GetId(),
				Version: 0, //Logic will update in Apply
				UpdatedAt: time.Now(),
			},
		}
		log.Printf("Event is created with type: %T", evt)
		return evt
	default:
		log.Println(v)
		return nil
	}
}
