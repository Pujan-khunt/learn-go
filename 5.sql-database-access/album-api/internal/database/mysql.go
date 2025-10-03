package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewConnection() (*sql.DB, error) {
	// Create a config object from environment variables.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "mysql"
	cfg.ParseTime = true // Apparently, its important for the Go's sql package to work correctly.

	var err error
	// Pass the config object after converting it to a connection string.
	// sql.Open will verify the driver availability and allocate memory for a sql.DB object.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("could not open sql connection: %w", err)
	}

	// Creates the actual connection to the mysql db using the connection string and the driver provided earlier
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping the database: %w", err)
	}

	return db, nil
}
