package rabbit

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func SetupTest() *AmqpDispatcher {
	d := NewDispatcher("amqp://guest:guest@localhost:5672/")

	return d

}

func TestAmqpDispatcher_Apply(t *testing.T) {

	dispatcher := SetupTest()
	// Command we are dispatching
	createsensor := CreateSensor{
		Name:       "living area",
		Sensortype: "temperature",
		CommandModel: CommandModel{
			Id: "2w3e-234w",
		},
	}
	//sensorcreated := `{"ID":"2w3e-234w","Version":0,"CreatedAt":"2021-10-07T18:01:23.7223548+03:00","UpdatedAt":"2021-10-07T18:01:23.7223548+03:00"}`
	// Event that will be constructed inside Apply and against which we will compare the actual result
	sensorcreated := SensorCreated{
		EventModel: EventModel{
			ID:      createsensor.Id,
			Version: 0,
		},
		Name:       createsensor.Name,
		Sensortype: createsensor.Sensortype,
	}
	// Apply Function is tested here
	err := dispatcher.Apply(context.Background(), createsensor)
	if err != nil {
		log.Fatal(err)
	}
	// Consume the message we sent above
	deliveries, _ := dispatcher.Channel.Consume(
		dispatcher.Queuename,
		"",
		true,
		true,
		false,
		true,
		nil)

	// Receive msg from channel and deserialize
	msg := <-deliveries
	var unpacked = SensorCreated{}
	err = json.Unmarshal(msg.Body, &unpacked)
	if err != nil {
		log.Fatal("Could not unpack payload from Apply Function")
	}
	// Assert that sensorcreated event expected is equal to unpacked event from function
	assert.Equal(t, sensorcreated.ID, unpacked.ID)
	assert.Equal(t, sensorcreated.Name, unpacked.Name)
	assert.Equal(t, sensorcreated.Sensortype, unpacked.Sensortype)

	// Close the connection
	_ = dispatcher.Channel.Close()

}
