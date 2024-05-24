package api

// ShortenUrlRequest represents a request to shorten a url
type ShortenUrlRequest struct {
	// URL to be shortened
	URL string `json:"url" binding:"url"`
}
