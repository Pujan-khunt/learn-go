package album

import "fmt"

// Service provides the business logic for the album operations.
type Service interface {
	CreateAlbum(album Album) (int64, error)
	GetAlbum(id int64) (*Album, error)
	GetAlbumsByArtist(artistName string) ([]Album, error)
}

// Implicitly implements the Service interface, by implementing all methods defined in the interface.
type albumService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &albumService{repo: repo}
}

// CreateAlbum implements Service.
func (a *albumService) CreateAlbum(album Album) (int64, error) {
	if album.Price < 0 {
		return 0, fmt.Errorf("price of album must be positive")
	}

	return a.repo.AddAlbum(album)
}

// GetAlbum implements Service.
func (a *albumService) GetAlbum(id int64) (*Album, error) {
	return a.repo.AlbumByID(id)
}

// GetAlbumsByArtist implements Service.
func (a *albumService) GetAlbumsByArtist(artistName string) ([]Album, error) {
	return a.repo.AlbumsByArtist(artistName)
}
