package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type DB struct {
	Connection *sql.DB
}

func NewDatabase() (*DB, error) {
	dataSource := os.Getenv("SQLITE_DB")
	dbDriver := os.Getenv("DB_DRIVER")

	conn, err := sql.Open(dbDriver, dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to establish the database connection with %v: %v", dataSource, err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to verify that the connection to database is still alive: %v", err)
	}

	return &DB{Connection: conn}, nil
}

func (db *DB) CloseDatabaseConnection() error {
	return db.Connection.Close()
}
