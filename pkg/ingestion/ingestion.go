package ingestion

import (
	"context"
	"errors"
	"fmt"

	pb_google_empty "github.com/golang/protobuf/ptypes/empty"
	pb "github.com/keelerh/omniscience/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type IngestionService struct {
	gdriveClient pb.GoogleDriveClient
}

func NewIngestionService() (*IngestionService, error) {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &IngestionService{
		gdriveClient: pb.NewGoogleDriveClient(cc),
	}, nil
}

func (s *IngestionService) Index(ctx context.Context, request *pb.IndexDocumentServiceRequest) (*pb_google_empty.Empty, error) {
	switch request.Service {
	case pb.DocumentService_GDRIVE:
		stream, err := s.gdriveClient.GetAll(context.Background(), &pb.GetAllDocumentsRequest{
			ModifiedSince: request.LastModifiedTime,
		})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to get all documents for Google Drive: %v", err))
		}
		resp, err := stream.Recv()
		fmt.Println("stream recv in ingestion: ", resp)
	default:
		stream, err := s.gdriveClient.GetAll(context.Background(), &pb.GetAllDocumentsRequest{
			ModifiedSince: request.LastModifiedTime,
		})
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Failed to get all documents for Google Drive: %v", err))
		}
		resp, err := stream.Recv()
		fmt.Println("stream recv in ingestion: ", resp)
	}

	return &pb_google_empty.Empty{}, nil
}
