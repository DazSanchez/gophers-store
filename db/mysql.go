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

// GetMySQLDBManager creates a Singleton instance of a IDBManager for MySQL database using the given config and returns a pointer to it.
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

// Connect opens and returns a connection to a MySQL database with the given mysql.Config as a DataSourceName.
// It panics if there's an error while connecting.
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
