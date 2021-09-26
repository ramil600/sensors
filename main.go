package main

import (
	//"context"
	"fmt"
	"github.com/ramil600/sensors/dbase"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	dbconn, err := dbase.New("root:root@tcp(127.0.0.1:3306)/sensors")

	if err != nil {
		fmt.Println("Could not find the database")
	}
	
	//dbconn.CreateTable(context.Background())
	sensor, err := dbconn.InsertSensor(1,"livingroom", "temp","heating")
	dbconn.GetSensor(sensor.Id)
}