package dbase

import (
	"log"
)

type Connection interface {
	InsertSensor(int, string, string, string) (Sensor, error)
	GetSensor(int) (Sensor, error)
}

//InsertSensor inserts new sensor type into sensortype table
func (dbc DbConn) InsertSensor(version int, name, sensortype, topic string) (Sensor, error) {
	s := Sensor{}
	if err := dbc.db.Ping(); err != nil {
		return s, err
	}

	stmt, err := dbc.db.Prepare("INSERT INTO sensortype(version, name, type, topic) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(version, name, sensortype, topic)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID= %d, rows affected = %d\n", lastID, rowCnt)
	s = Sensor{
		Id:         int(lastID),
		Version:    int(version),
		Name:       name,
		Sensortype: sensortype,
	}

	return s, nil

}

//GetSensor get sensor from sensor types with Id
func (dbc DbConn) GetSensor(id int) (Sensor, error) {
	s := Sensor{}
	if err := dbc.db.Ping(); err != nil {
		return s, err
	}
	err := dbc.db.QueryRow("SELECT version, name, type, topic FROM sensortype WHERE id = ?", id).
		Scan(&s.Version, &s.Name, &s.Sensortype, &s.Topic)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s.Version, s.Name, s.Sensortype, s.Topic)
	return s, nil
}
