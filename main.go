package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramil600/sensors/config"
	"github.com/ramil600/sensors/rabbit"
)

func main() {

	cfg := config.NewConfig()
	dispatcher := rabbit.NewDispatcher(cfg.PublishDSN)

	dispatcher.Subscribe(cfg.EventsQueue)

	dispatcher.Shutdown()

	//cfg := config.NewConfig()
	//
	//dbconn, err := dbase.New(dbase.MYSQL_DSN)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
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
