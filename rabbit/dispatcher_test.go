package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func SetupTest() (*AmqpDispatcher, func()) {
	d := NewDispatcher("amqp://guest:guest@localhost:5672/")
	teardown := func() {
		d.Conn.Close()
		fmt.Println("Closing connection")
	}
	return d, teardown

}

func TestAmqpDispatcher_Apply(t *testing.T) {

	dispatcher, teardown := SetupTest()
	//Close the connection to dispatcher when test is done
	defer teardown()
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
		log.Fatal("TestAmqpDispatcher#Apply: Could not unpack payload from Apply Function")
	}
	// Assert that sensorcreated event expected is equal to unpacked event from function
	assert.Equal(t, sensorcreated.ID, unpacked.ID)
	assert.Equal(t, sensorcreated.Name, unpacked.Name)
	assert.Equal(t, sensorcreated.Sensortype, unpacked.Sensortype)

}
func TestEventFromCommand_UpdateSensor(t *testing.T) {

	com := UpdateSensor{
		Name: "living room",
		Sensortype: "tempreature",
		CommandModel: CommandModel{
			Id: "u313-ds34-su3",
			Type: "data",
		},
	}
	expected :=
		SensorUpdated{
		Name: "living room",
		Sensortype: "temperature",
		EventModel: EventModel{
			ID: "u313-ds34-su3",
			Version: 0,
		},

		}
	actual := EventFromCommand(com)

	assert.Equal(t,expected.ID, actual.GetId())

}
