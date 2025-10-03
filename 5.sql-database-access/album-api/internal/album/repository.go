package album

import (
	"database/sql"
	"fmt"
)

// Repository handles all the database interactions for albums.
// We use an interface to allow for easy mocking in tests.
type Repository interface {
	AddAlbum(album Album) (int64, error)
	AlbumByID(id int64) (*Album, error)
	AlbumsByArtist(artistName string) ([]Album, error)
}

// mySQLRepository implements the Repository interface for a MySQL database.
type mySQLRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &mySQLRepository{db: db}
}

// AlbumByID Returns the album from the database with a given id.
func (r *mySQLRepository) AlbumByID(id int64) (*Album, error) {
	var album Album

	// Since we are only expecting a single row as a response, we use the QueryRow method
	// QueryRow doesn't return an error and always returns a non-nil value.
	row := r.db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	// QueryRow waits until the user uses the row.Scan method which will throw the error(if any)
	// which was supposed to be returned by the QueryRow function.
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		// This error(if any) is returned by the QueryRow function.
		// Checked error for query returning zero rows.
		if err == sql.ErrNoRows {
			return &album, fmt.Errorf("albumById %d: no such album", id)
		}
		// Unchecked error
		return &album, fmt.Errorf("albumById: %d: %v", id, err)
	}
	return &album, nil
}

// AlbumsByArtist Returns all the albums with a given artist name
func (r *mySQLRepository) AlbumsByArtist(artistName string) ([]Album, error) {
	// Album slice to hold data from returned rows.
	var albums []Album

	// Run select query on DB to get albums with a specified artist.
	rows, err := r.db.Query("SELECT * FROM album WHERE artist = ?", artistName)
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

// AddAlbum Inserts a new album into the database.
func (r *mySQLRepository) AddAlbum(album Album) (int64, error) {
	// Exec() is used to run queries which don't return any rows.
	result, err := r.db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
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
