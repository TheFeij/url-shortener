package store

// SaveUrlRequest represents a store request to save an url record into
// the urls table in the database
type SaveUrlRequest struct {
	// OriginalUrl client's original url (long url)
	originalUrl string
	// ShortUrl shortened url returned to the client
	shortUrl string
}

// NewSaveUrlRequest return a new SaveUrlRequest
func NewSaveUrlRequest(originalUrl string, shortUrl string) *SaveUrlRequest {
	return &SaveUrlRequest{originalUrl: originalUrl, shortUrl: shortUrl}
}

// OriginalUrl getter for originalUrl
func (s SaveUrlRequest) OriginalUrl() string {
	return s.originalUrl
}

// ShortUrl getter for shortUrl
func (s SaveUrlRequest) ShortUrl() string {
	return s.shortUrl
}
