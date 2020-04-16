package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	AllMachines() ([]*MachineSchema, error)
	//NewMachine() (*Machine, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Print("Database connected")

	return &DB{db}, nil
}
