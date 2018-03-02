package ingestion

import (
	"time"

	pb "github.com/keelerh/omniscience/protos"
	"github.com/keelerh/omniscience/server/vendor/github.com/olivere/elastic"
)

// Document is a structure used for serializing/deserializing data in Elasticsearch.
type Document struct {
	Id           string                `json:"id"`
	Name         string                `json:"name,omitempty"`
	Description  string                `json:"description,omitempty"`
	Service      pb.DocumentService    `json:"image,omitempty"`
	Content      string                `json:"created,omitempty"`
	Url          string                `json:"created,omitempty"`
	LastModified time.Time             `json:"created,omitempty"`
	Suggest      *elastic.SuggestField `json:"suggest_field,omitempty"`
}
