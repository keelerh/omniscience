package document_fetcher

import "time"

type DocumentFetcher interface {
	// Fetches documents modified since a set time.
	Fetch(modifiedSince time.Time) error
}