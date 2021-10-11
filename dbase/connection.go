package dbase

import (
	"log"
)

const MYSQL_DSN = "root:root@tcp(127.0.0.1:3306)/sensors"

type Connection interface {
	InsertSensor(int, string, string, string) (Sensor, error)
	GetSensor(int) (Sensor, error)
	UpdateSensor (int, int, string, string, string) (Sensor, error)
}

//InsertSensor inserts new sensor type into sensortype table returns constructed Sensor struct
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

	log.Printf("Sensor inserted with ID= %d, rows affected = %d\n", lastID, rowCnt)
	s = Sensor{
		Id:         int(lastID),
		Version:    int(version),
		Name:       name,
		Sensortype: sensortype,
		Topic: topic,
	}

	return s, nil

}

// UpdateSensor changes version, sensortype or topic of the sensor returns updated Sensor struct
func (dbc DbConn) UpdateSensor(id int, version int, name, sensortype, topic string) (Sensor, error) {
	stmt, err := dbc.db.Prepare("UPDATE sensortype SET version=?, name=?, type=?, topic=?  WHERE id=?")
	if err != nil{
		log.Fatal(err)
	}
	defer stmt.Close()


	res, err := stmt.Exec(version, name, sensortype, topic, id)
	if err != nil {
		log.Fatal(err)
	}
	CountId, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Number of rows affected: ", CountId)

	return Sensor{Id: id}, nil

}

//GetSensor returns Sensor struct matched by Id
func (dbc DbConn) GetSensor(id int) (Sensor, error) {
	s := Sensor{}

	err := dbc.db.QueryRow("SELECT version, name, type, topic FROM sensortype WHERE id = ?", id).
		Scan(&s.Version, &s.Name, &s.Sensortype, &s.Topic)
	if err != nil {
		return s, err
	}
	log.Println(s.Version, s.Name, s.Sensortype, s.Topic)
	return s, nil
}
