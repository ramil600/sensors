package dbase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DbLogger struct {
	db *sql.DB
}


func New(dsn string) (*DbLogger, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DbLogger{db:db}, err

}

func (m DbLogger)CreateTable(ctx context.Context){

	//db, err := sql.Open("mysql", "root:root@/sensors")
/*
	if err != nil {
		fmt.Println("Could not find the database")
	}

	defer db.Close()
	
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
*/	
    query := `CREATE TABLE IF NOT EXISTS
	sensortype(
		id int primary key auto_increment,
		version int,
		name varchar(20),
		type varchar(20),
		topic varchar(30) unique
	)`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := m.db.ExecContext(ctx, query)

	if err != nil {
		log.Printf("Error is %v\n", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	} 
	log.Println(rows, "rows affected")

}