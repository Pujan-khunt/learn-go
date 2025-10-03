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

func AddAlbum(album Album) (int64, error) {
	// Exec() is used to run queries which don't return any rows.
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	// Get the ID of the insertion to return to the caller.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func AlbumByID(id int) (Album, error) {
	var album Album

	// Since we are only expecting a single row as a response, we use the QueryRow method
	// QueryRow doesn't return an error and always returns a non-nil value.
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	// QueryRow waits until the user uses the row.Scan method which will throw the error(if any)
	// which was supposed to be returned by the QueryRow function.
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		// This error(if any) is returned by the QueryRow function.
		// Checked error for query returning zero rows.
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumById %d: no such album", id)
		}
		// Unchecked error
		return album, fmt.Errorf("albumById: %d: %v", id, err)
	}
	return album, nil
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

	// Get all albums of this artist.
	artistName := "John Coltrane"
	albums, err := AlbumsByArtist(artistName)
	if err != nil {
		log.Fatalf("Error fetching album of the artist: %q. Error: %v\n", artistName, err)
	}
	log.Printf("Albums Found: %v\n", albums)

	// Get the album by this id.
	id := 2
	album, err := AlbumByID(id)
	if err != nil {
		log.Fatalf("Error fetching album with id: %d. Error: %v\n", id, err)
	}
	log.Printf("Album found: %v\n", album)

	// Add a new album with these values
	albumID, err := AddAlbum(Album{
		Title:  "Sajna",
		Artist: "Pujan Khunt",
		Price:  200,
	})
	if err != nil {
		log.Fatalf("Error adding album: %v", err)
	}
	log.Printf("Album added with id: %d", albumID)
}
