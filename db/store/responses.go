package store

// SaveUrlResponse represents the response to a SaveUrlRequest
type SaveUrlResponse struct {
	// originalUrl client's original url (long url)
	originalUrl string
	// shortUrl shortened url returned to the client
	shortUrl string
}

// NewSaveUrlResponse return a new SaveUrlResponse
func NewSaveUrlResponse(originalUrl string, shortUrl string) *SaveUrlResponse {
	return &SaveUrlResponse{originalUrl: originalUrl, shortUrl: shortUrl}
}

// OriginalUrl getter for originalUrl
func (s SaveUrlResponse) OriginalUrl() string {
	return s.originalUrl
}

// ShortUrl getter for shortUrl
func (s SaveUrlResponse) ShortUrl() string {
	return s.shortUrl
}
