package ingestion

import (
	"context"
	"fmt"
	"io"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/olivere/elastic"
	"github.com/keelerh/omniscience/pkg/elasticsearch"
)

const (
	defaultDescriptionLengthInChars = 400
)

type Ingester struct {
	elasticClient *elastic.Client
}

func NewIngester(esClient *elastic.Client) (*Ingester, error) {
	return &Ingester{
		elasticClient: esClient,
	}, nil
}

func (s *Ingester) Ingest(stream pb.Ingester_IngestServer) error {
	for {
		doc, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if err := s.index(doc); err != nil {
			return err
		}
	}

	return nil
}

func (s *Ingester) index(d *pb.Document) error {
	if err := s.createIndexIfNotExists(context.Background()); err != nil {
		return err
	}

	if err := s.indexDocument(context.Background(), d); err != nil {
		return err
	}

	// Flush to make sure the documents got written.
	_, err := s.elasticClient.Flush().Index(elasticsearch.Index).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s *Ingester) createIndexIfNotExists(ctx context.Context) error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := s.elasticClient.IndexExists(elasticsearch.Index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, err := s.elasticClient.CreateIndex(elasticsearch.Index).BodyString(elasticsearch.Mapping).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged.
		}
	}

	return nil
}

func (s *Ingester) indexDocument(ctx context.Context, d *pb.Document) error {
	// Index a document (using JSON serialization).
	id := d.GetId().GetId()
	lastModifiedTime := d.LastModified
	lastModified, err := ptypes.Timestamp(lastModifiedTime)
	if err != nil {
		return err
	}
	description := d.Description
	if description == "" {
		description = d.Content
		if len(description) > defaultDescriptionLengthInChars {
			description = description[:defaultDescriptionLengthInChars] + "..."
		}
	}
	doc := Document{
		Id:           id,
		Title:        d.Title,
		Description:  description,
		Service:      d.Service,
		Content:      d.Content,
		Url:          d.Url,
		LastModified: lastModified,
	}

	put, err := s.elasticClient.Index().
		Index(elasticsearch.Index).
		Type("_doc").
		Id(id).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("Indexed document %s with ID %s to index %s\n", doc.Title, put.Id, put.Index)
	return nil
}
