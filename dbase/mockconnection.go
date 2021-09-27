package dbase

import (
	"github.com/stretchr/testify/mock"
)

type MyTestDB struct {
	// add a Mock object instance
	mock.Mock

	// other fields go here as normal
}

func (dbc MyTestDB) InsertSensor(version int, name, sensortype, topic string) (Sensor, error) {
	args := dbc.Called()

	return args.Get(0).(Sensor), nil

}

func (dbc MyTestDB) GetSensor(id int) (Sensor, error) {
	args := dbc.Called()
	return args.Get(0).(Sensor), nil
}
