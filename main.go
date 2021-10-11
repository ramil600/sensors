package main

import (
	//"github.com/ramil600/sensors/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("hello")
    //cfg := config.NewConfig()
	//
	//dbconn, err := dbase.New(dbase.MYSQL_DSN)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//sensor, err := dbconn.InsertSensor(0,"livingroom","temperature", "dispatcher")
	//if err != nil {
	//	log.Fatal(err)
	//
	//}
	//
	//
	//
	//fmt.Println("Sensor was created with id: ", sensor.Id)
	//
	//
	//
	//
	//
	//sensor, err = dbconn.UpdateSensor(sensor.Id, 100, sensor.Name,sensor.Sensortype, sensor.Topic)
	//
	//log.Println(sensor.Id)
	//

	/*
	dpt := rabbit.NewDispatcher(cfg.PublishDSN)
	ctx := context.Background()
	cmdCreateSensor := rabbit.CreateSensor{
		CommandModel: rabbit.CommandModel{
			Id: "w3wq-2da2",
		},
		Sensortype: "temp_sensor",
		Name:       "livingroom01",
	}

	err := dpt.Apply(ctx, cmdCreateSensor)

	if err != nil {
		fmt.Println("Could Apply Command on Create Sensor")
	}
    */
}
