package service

import (
	"github.com/ramil600/sensors/dbase"
)

type MyService struct {
	db dbase.Connection
}

func NewService(mydb dbase.Connection) (MyService, error) {
	return MyService{db: mydb}, nil
}

func (s MyService) CreateSensor(version int, name, sensortype, topic string) {
	s.db.InsertSensor(version, name, sensortype, topic)
}
