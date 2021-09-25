package main

import (
	"context"
	"fmt"
	"github.com/ramil600/sensors/dbase"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	dblogger, err := dbase.New("root:root@/sensors")

	if err != nil {
		fmt.Println("Could not find the database")
	}
	
	/*
	dblogger.db.SetConnMaxLifetime(time.Minute * 3)
	dblogger.db.SetMaxOpenConns(10)
	dblogger.db.SetMaxIdleConns(10)
	*/

	dblogger.CreateTable(context.Background())


}