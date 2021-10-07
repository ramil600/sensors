package rabbit

import (
	"log"
	"time"
)

type Event interface {
	GetId() string
	EventVersion() int
	EventUpdatedAt() time.Time
}

type EventModel struct {
	ID        string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m EventModel) GetId() string {
	return m.ID
}
func (m EventModel) EventVersion() int {
	return m.Version
}

func (m EventModel) EventUpdatedAt() time.Time {
	return m.UpdatedAt
}

type SensorCreated struct {
	EventModel
	Name       string
	Sensortype string
}

type WarnEvent struct {
	EventModel
}

type ErrorEvent struct {
	EventModel
}

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
	default:
		log.Println(v)
		return nil
	}
}
