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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGoogleDriveClient(conn)

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
	r, err := c.GetAll(context.Background(), &pb.GetAllDocumentsRequest{
		ModifiedSince: modifiedSince,
	})
	if err != nil {
		log.Fatalf("Failed to get all documents: %v", err)
	}
	log.Printf("Documents: %s", r.Documents)
}
