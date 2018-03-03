package ingestion

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/golang/protobuf/ptypes"
	pb_google_empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/olivere/elastic"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	defaultDescriptionLengthInChars = 120
)

type IngestionService struct {
	elasticClient *elastic.Client
	gdriveClient  pb.GoogleDriveClient
}

func NewIngestionService(esClient *elastic.Client) (*IngestionService, error) {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &IngestionService{
		elasticClient: esClient,
		gdriveClient:  pb.NewGoogleDriveClient(cc),
	}, nil
}

func (s *IngestionService) Index(ctx context.Context, request *pb.IndexDocumentServiceRequest) (*pb_google_empty.Empty, error) {
	switch request.Service {
	case pb.DocumentService_GDRIVE:
		stream, err := s.gdriveClient.GetAll(ctx, &pb.GetAllDocumentsRequest{
			ModifiedSince: request.LastModified,
		})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to get documents for Google Drive: %v", err))
		}
		for {
			doc, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err := s.index(ctx, doc); err != nil {
				return nil, err
			}
		}
	}

	return &pb_google_empty.Empty{}, nil
}

func (s *IngestionService) index(ctx context.Context, d *pb.Document) error {
	if err := s.createIndexIfNotExists(ctx); err != nil {
		return err
	}

	if err := s.indexDocument(ctx, d); err != nil {
		return err
	}

	// Flush to make sure the documents got written.
	_, err := s.elasticClient.Flush().Index(index).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *IngestionService) createIndexIfNotExists(ctx context.Context) error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := s.elasticClient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		// Create a new index.
		createIndex, err := s.elasticClient.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged.
		}
	}

	return nil
}

func (s *IngestionService) indexDocument(ctx context.Context, d *pb.Document) error {
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
			description = description[:defaultDescriptionLengthInChars]
		}
	}
	fmt.Println(d.Url)
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
		Index(index).
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
