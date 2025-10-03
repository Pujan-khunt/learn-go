package album

import "log"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetAlbumsByArtist Simulates an API call to get albums by an artist.
func (h *Handler) GetAlbumsByArtist(artistName string) {
	log.Printf("HANDLER: fetching albums for artist: %q", artistName)
	albums, err := h.service.GetAlbumsByArtist(artistName)
	if err != nil {
		log.Fatalf("HANDLER ERROR: %v", err)
	}
	log.Printf("HANDLER SUCCESS: Albums Found: %v", albums)
}

// GetAlbumByID Simulates an API call to get a single album.
func (h *Handler) GetAlbumByID(id int64) {
	log.Printf("HANDLER: fetching album with ID: %d", id)
	album, err := h.service.GetAlbum(id)
	if err != nil {
		log.Fatalf("HANDLER ERROR: %v", err)
	}
	log.Printf("HANDLER SUCCESS: Album Found: %v", album)
}

// AddNewAlbum Simulates an API call to add a new album.
func (h *Handler) AddNewAlbum(title, artist string, price float32) {
	log.Printf("HANDLER: adding new album")
	album := Album{
		Title:  title,
		Artist: artist,
		Price:  price,
	}
	albumID, err := h.service.CreateAlbum(album)
	if err != nil {
		log.Fatalf("HANDLER ERROR: %v", err)
	}
	log.Printf("HANDLER SUCCESS: Album added with ID: %d", albumID)
}
