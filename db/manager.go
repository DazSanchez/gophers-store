package db

import (
	"database/sql"
	"log"
)

var Instance *sql.DB

type IDBManager interface {
	Connect() (*sql.DB, error)
}

func Init(manager IDBManager) {
	log.Println("Init DB connection")

	db, err := manager.Connect()
	if err != nil {
		panic("can't connect to db: " + err.Error())
	}

	Instance = db
}

func Close() {
	if Instance != nil {
		log.Println("Close DB connection")
		Instance.Close()
	}
}
