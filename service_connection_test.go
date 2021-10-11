package main

import (
	"github.com/ramil600/sensors/dbase"
	"github.com/ramil600/sensors/service"
	"log"
	"testing"
)

//TestInsertSensor creates service and mock connection to test without db

func TestInsertSensor(t *testing.T) {
	mockConn := new(dbase.MyTestDB)
	mockConn.On("InsertSensor").Return(dbase.Sensor{Name: "name", Topic: "topic"}, nil)

	mytestservice, err := service.NewSensorService(mockConn)
	if err != nil {
		log.Fatal(err)
	}
	//mytestservice doesn't know which connection it uses as it deals with Connection interface
	mytestservice.CreateSensor(1, "name", "", "")
	mockConn.AssertExpectations(t)
}
