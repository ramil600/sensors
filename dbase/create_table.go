package dbase

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DbConn struct {
	db *sql.DB
}

//Return new DbConn struct with set parameters
func New(dsn string) (*DbConn, error) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DbConn{db:db}, err

}

// CreateTable only needs to be run to create a table if it doesnt exist for sensor types
func (m DbConn)CreateTable(ctx context.Context){
	
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