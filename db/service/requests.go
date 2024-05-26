package service

// SaveUrlRequest represents a service request to save an url record into
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

// GetOriginalUrlRequest represents a service request to get the original url
// relative to the short url
type GetOriginalUrlRequest struct {
	// ShortUrl shortened url
	shortUrl string
}

// NewGetOriginalUrlRequest return a new GetOriginalUrlRequest
func NewGetOriginalUrlRequest(shortUrl string) *GetOriginalUrlRequest {
	return &GetOriginalUrlRequest{shortUrl: shortUrl}
}

// ShortUrl getter for shortUrl
func (g GetOriginalUrlRequest) ShortUrl() string {
	return g.shortUrl
}
