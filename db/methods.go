package db

import (
	"url-shortener/db/models"
	"url-shortener/db/service"
)

// SaveUrl creates a models.Url instance from (req SaveUrlRequest) and saves it into the database
// returns an error instance if any
func (d database) SaveUrl(req *service.SaveUrlRequest) (*service.SaveUrlResponse, error) {
	// create model instance
	record := models.Url{
		OriginalUrl: req.OriginalUrl(),
		ShortUrl:    req.ShortUrl(),
	}

	// save model instance to the database
	if err := d.db.Create(&record).Error; err != nil {
		return nil, err
	}

	// create a response instance
	response := service.NewSaveUrlResponse(
		record.OriginalUrl,
		record.ShortUrl,
	)

	// return results
	return response, nil
}
