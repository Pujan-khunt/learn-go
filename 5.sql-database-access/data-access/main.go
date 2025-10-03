package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Represents the DB handle.
var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func AlbumsByArtist(artistName string) ([]Album, error) {
	// Album slice to hold data from returned rows.
	var albums []Album

	// Run select query on DB to get albums with a specified artist.
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artistName)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
	}
	defer rows.Close()

	// Loop through returned rows to convert data into the strongly typed object.
	for rows.Next() {
		var album Album
		if rowScanErr := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); rowScanErr != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, rowScanErr)
		}
		albums = append(albums, album)
	}

	// rows.Err returns an error (if any) indicating that the rows.Next() was terminated due to rows exhaustion or an error was occured.
	// Important to check rows.Err() after looping through all rows
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist: %q: %v", artistName, err)
	}

	return albums, nil
}

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
	// sql.Open will verify the driver availability and allocate memory for a sql.DB object.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error while connecting to MySQL DB. ", err)
	}

	// Creates the actual connection to the mysql db using the connection string and the driver provided earlier
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Error checking connection with DB. ", err)
	}

	log.Println("MySQL DB connected and ready for operation.")

	artistName := "John Coltrane"
	albums, err := AlbumsByArtist(artistName)
	if err != nil {
		log.Fatalf("Error fetching albume of the artist: %q. Error: %v\n", artistName, err)
	}

	log.Printf("Albums Found: %v\n", albums)
}
