package document_fetcher

import "time"

type DocumentFetcher interface {
	// Fetches documents modified since a set time and forwards them to the IngestionService.
	Fetch(modifiedSince time.Time) error
}
