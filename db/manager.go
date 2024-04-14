package db

import (
	"context"
	"database/sql"
	"log"
)

// Instance is a Singleton that holds a reference to a opened database connection.
var Instance *sql.DB

// IDBManager is a generic interface Database conectors can implement to abstract connection details.
type IDBManager interface {
	// Connect initialize a new database connection and saves it as a Singleton.
	Connect() (*sql.DB, error)
}

// Init creates a new database connection instance using the given IDBManager
func Init(manager IDBManager) {
	log.Println("Init DB connection")

	db, err := manager.Connect()
	if err != nil {
		panic("can't connect to db: " + err.Error())
	}

	Instance = db
}

// Close terminates the current database Singleton connection if any.
func Close() {
	if Instance != nil {
		log.Println("Close DB connection")
		Instance.Close()
	}
}

// TxBegin creates a Transaction context to run queries upon.
// It uses a Background context instance to create the Transaction.
// It panics if can't create the Transaction.
func TxBegin() (*sql.Tx, error) {
	return Instance.BeginTx(context.Background(), nil)
}
