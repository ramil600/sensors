package main

import (
	"context"
	"fmt"
	"github.com/ramil600/sensors/rabbit"
	//"github.com/ramil600/sensors/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//dbconn, err := dbase.New(dbase.MYSQL_DSN)
	dpt := rabbit.NewDispatcher()
	ctx := context.Background()
	cmdCreateSensor := rabbit.CreateSensor{
		CommandModel: rabbit.CommandModel{
			Id: "w3wq-2da2",
		},
		Sensortype: "temp_sensor",
		Name: "livingroom01",
	}


	err := dpt.Apply(ctx,cmdCreateSensor)



	if err != nil {
		fmt.Println("Could Apply Command on Create Sensor")
	}



}
