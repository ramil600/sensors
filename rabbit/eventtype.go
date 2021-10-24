package rabbit

import (
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

type SensorUpdated struct {
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


