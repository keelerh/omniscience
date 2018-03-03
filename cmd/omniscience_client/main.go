package main

import (
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/keelerh/omniscience/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address                       = "localhost:50051"
	defaultModifiedSinceTimestamp = "02 Jan 06 15:04 MST"
)

func main() {
	timestamp := defaultModifiedSinceTimestamp
	if len(os.Args) > 1 {
		timestamp = os.Args[1]
	}
	modifiedSinceTimestamp, err := time.Parse(time.RFC822, timestamp)
	if err != nil {
		log.Fatalf("unable to parse modified since timestamp: %v", err)
	}
	modifiedSince, err := ptypes.TimestampProto(modifiedSinceTimestamp)
	if err != nil {
		log.Fatalf("unable to parse modified since timestamp as proto: %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ingestionClient := pb.NewIngestionClient(conn)
	_, err = ingestionClient.Index(context.Background(), &pb.IndexDocumentServiceRequest{
		Service:      pb.DocumentService_GDRIVE,
		LastModified: modifiedSince,
	})
	if err != nil {
		log.Fatalf("failed to index documents for Google Drive: %v", err)
	}
}
