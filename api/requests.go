package api

// ShortenUrlRequest represents a request to shorten a url
type ShortenUrlRequest struct {
	// URL to be shortened
	URL string `json:"url" binding:"url"`
}

// RedirectShortUrl represents a request to be redirected to the original url
type RedirectShortUrl struct {
	// ShortUrl short url
	ShortUrl string `uri:"short_url" binding:"required,len=6"`
}
