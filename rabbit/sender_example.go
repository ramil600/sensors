package rabbit

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func Sendmyqueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", //queue name
		false,   //durable
		false,   //AutoDelete
		false,   //exclusive
		false,   //noWait
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {

		body0 := WarnEvent{
			EventModel{ID: strconv.Itoa(i),
				CreatedAt: time.Now()},
		}
		payload, _ := json.Marshal(&body0)

		err = ch.Publish(
			"",     // exchange
			q.Name, //routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				Body: payload,
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Print(" [x] sent payload 100 times\n")

}
