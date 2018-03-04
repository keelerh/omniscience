package ingestion

import (
	"time"

	"github.com/olivere/elastic"
)

// Document is a structure used for serializing/deserializing data in Elasticsearch.
type Document struct {
	Id           string                `json:"id"`
	Title        string                `json:"title,omitempty"`
	Description  string                `json:"description,omitempty"`
	Service      string                `json:"image,omitempty"`
	Content      string                `json:"created,omitempty"`
	Url          string                `json:"created,omitempty"`
	LastModified time.Time             `json:"created,omitempty"`
	Suggest      *elastic.SuggestField `json:"suggest_field,omitempty"`
}
