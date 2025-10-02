package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Represents the DB handle.
var db *sql.DB

func main() {
	// Create a config object and pass relevant values.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "mysql"

	var err error
	// Pass the config object after converting it to a connection string.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error while connecting to MySQL DB. ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Error checking connection with DB. ", err)
	}

	log.Println("MySQL DB connected and ready for operation.")
}
