package main

import (
	"log"
	"os"
	"time"

	"github.com/keelerh/omniscience/server/vendor/github.com/golang/protobuf/ptypes"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/keelerh/omniscience/server/vendor/golang.org/x/net/context"
	"github.com/keelerh/omniscience/server/vendor/google.golang.org/grpc"
)

const (
	address                       = "localhost:50051"
	defaultModifiedSinceTimestamp = "02 Jan 06 15:04 MST"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	timestamp := defaultModifiedSinceTimestamp
	if len(os.Args) > 1 {
		timestamp = os.Args[1]
	}
	modifiedSinceTimestamp, err := time.Parse(time.RFC822, timestamp)
	if err != nil {
		log.Fatalf("Unable to parse modified since timestamp: %v", err)
	}
	modifiedSince, err := ptypes.TimestampProto(modifiedSinceTimestamp)
	if err != nil {
		log.Fatalf("Unable to parse modified since timestamp as proto: %v", err)
	}

	ingestionClient := pb.NewIngestionClient(conn)
	_, err = ingestionClient.Index(context.Background(), &pb.IndexDocumentServiceRequest{
		Service:      pb.DocumentService_GDRIVE,
		LastModified: modifiedSince,
	})
	if err != nil {
		log.Fatalf("Failed to index documents for Google Drive: %v", err)
	}
}
