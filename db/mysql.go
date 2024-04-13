package db

import (
	"database/sql"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	lock         = &sync.Mutex{}
	mysqlManager *MySQLDBManager
)

type MySQLDBManager struct {
	dataSourceName string
}

func GetMySQLDBManager(config mysql.Config) *MySQLDBManager {
	if mysqlManager != nil {
		return mysqlManager
	}

	lock.Lock()
	defer lock.Unlock()

	if mysqlManager == nil {
		mysqlManager = &MySQLDBManager{
			dataSourceName: config.FormatDSN(),
		}
	}

	return mysqlManager
}

func (dbm MySQLDBManager) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbm.dataSourceName)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
