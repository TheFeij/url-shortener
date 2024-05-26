package service

// DBService defines methods to interact with the database
type DBService interface {
	// SaveUrl creates a models.Url instance from (req SaveUrlRequest) and saves it into the database
	// returns an error if any
	SaveUrl(req *SaveUrlRequest) (*SaveUrlResponse, error)
	// GetOriginalUrl returns the original url of a shortened url from the database
	// returns an error if any
	GetOriginalUrl(req *GetOriginalUrlRequest) (*GetOriginalUrlResponse, error)
}
