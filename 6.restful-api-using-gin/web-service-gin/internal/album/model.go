package album

// Encoding each fields with json encoder and ordering it to keep the key names as indicated in the double quotes.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
