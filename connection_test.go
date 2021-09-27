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

	myservice, err := service.NewService(mockConn)
	if err != nil {
		log.Fatal(err)
	}
	//myservice doesn't know which connection it uses as it deals with Connection interface
	myservice.CreateSensor(1, "name", "", "")
	mockConn.AssertExpectations(t)
}
