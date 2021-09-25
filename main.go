package main

import (
	"context"
	"fmt"
	"github.com/ramil600/sensors/dbase"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	dbconn, err := dbase.New("root:root@/sensors")

	if err != nil {
		fmt.Println("Could not find the database")
	}
	
	dbconn.CreateTable(context.Background())

}