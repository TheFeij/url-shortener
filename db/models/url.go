package models

// Url defines the model for the urls table.
// It represents a URL with its shortened version and the original version.
type Url struct {
	// The shortened version of the URL, primary key
	ShortUrl string `gorm:"column:short_url;primaryKey;type:varchar"`
	// The original URL
	OriginalUrl string `gorm:"column:original_url;type:varchar;not null"`
}
