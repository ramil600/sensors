package rabbit

import (

	"github.com/streadway/amqp"
	"log"
)

func Processmyqueue() {
	Sendmyqueue()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	consumech, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	wait := make(chan interface{})

	go func() {
		for msg := range consumech {
			log.Printf("Received message: %s\n", msg.Body)
		}
	}()

	log.Println("Wait for messages. To exit press CTRL+C\n")

	<-wait
}
