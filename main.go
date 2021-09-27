package main

import (
	"fmt"
	"github.com/ramil600/sensors/dbase"
	//"github.com/ramil600/sensors/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbconn, err := dbase.New(dbase.MYSQL_DSN)

	if err != nil {
		fmt.Println("Could not find the database")
	}

	//dbconn.CreateTable(context.Background())
	//sensor, err := dbconn.InsertSensor(1,"livingroom", "temp","heating")
	//dbconn.GetSensor(sensor.Id)
	dbconn.GetSensor(1)

}
