package store

// DBService defines methods to interact with the database
type DBService interface {
	// SaveUrl creates a models.Url instance from (req SaveUrlRequest) and saves it into the database
	// returns an error instance if any
	SaveUrl(req *SaveUrlRequest) (*SaveUrlResponse, error)
}
