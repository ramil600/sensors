package service

import (
	"github.com/ramil600/sensors/dbase"
)

type SensorService struct {
	db dbase.Connection
}

func NewSensorService(mydb dbase.Connection) (SensorService, error) {
	return SensorService{db: mydb}, nil
}

func (s SensorService) CreateSensor(version int, name, sensortype, topic string) {
	s.db.InsertSensor(version, name, sensortype, topic)
}
