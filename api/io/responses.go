package io

// ShortenUrlResponse represents a response to shorten url
type ShortenUrlResponse struct {
	// OriginalUrl client's original url (long url)
	OriginalUrl string `json:"original_url"`
	// ShortUrl shortened url returned to the client
	ShortUrl string `json:"short_url"`
}
