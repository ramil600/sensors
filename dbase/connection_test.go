package dbase

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}

	return db, mock
}

var sensor = Sensor{
	Id:         0,
	Version:    0,
	Name:       "living room",
	Sensortype: "temperature",
	Topic:      "dispatcher",
}

func TestDbConn_GetSensor(t *testing.T) {
	db, mock := NewMock()

	dbcon := DbConn{
		db: db,
	}

	query := "SELECT version, name, type, topic FROM sensortype WHERE id = ?"

	rows := sqlmock.NewRows([]string{"version", "name", "type", "topic"}).
		AddRow(sensor.Version, sensor.Name, sensor.Sensortype, sensor.Topic)
	mock.ExpectQuery(query).WithArgs(sensor.Id).WillReturnRows(rows)

	returned, err := dbcon.GetSensor(sensor.Id)
	assert.Equal(t, sensor, returned)
	assert.NoError(t, err)
}

func TestDbConn_InsertSensor(t *testing.T) {
	db, mock := NewMock()
	dbcon := DbConn{
		db: db,
	}
	query := "INSERT INTO sensortype\\(version, name, type, topic\\) VALUES\\(\\?,\\?,\\?,\\?\\)"



	mock.ExpectPrepare(query).ExpectExec().WithArgs(sensor.Version,sensor.Name, sensor.Sensortype,sensor.Topic).
		WillReturnResult(sqlmock.NewResult(0,1))

	actual, err := dbcon.InsertSensor(sensor.Version, sensor.Name,sensor.Sensortype, sensor.Topic)
	assert.Equal(t, sensor,actual)
	assert.NoError(t, err)



}
