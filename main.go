package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ramil600/sensors/config"
	"github.com/ramil600/sensors/rabbit"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {


	cfg := config.NewConfig()
	dispatcher := rabbit.NewDispatcher(cfg.PublishDSN)

	dispatcher.Subscribe(cfg.EventsQueue)
<<<<<<< HEAD

	dispatcher.Shutdown()
=======
>>>>>>> 85a864731c6d9a760fad756ac9928f3d5922d315

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case sht := <-shutdown:
		dispatcher.Shutdown()
		log.Println("main(): Dispatcher is shutting down: ", sht)
	}
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
