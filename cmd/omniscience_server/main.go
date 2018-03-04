// protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

package main

import (
	"log"
	"net"

	"github.com/keelerh/omniscience/pkg/ingestion"
	pb "github.com/keelerh/omniscience/protos"
	"google.golang.org/grpc"
	"github.com/olivere/elastic"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Obtain a client and connect to the default Elasticsearch installation on 127.0.0.1:9200.
	esClient, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("failed to initialise Elasticsearch client: %v", err)
	}

	ingester, err := ingestion.NewIngester(esClient)
	if err != nil {
		log.Fatalf("failed to instantiate Ingester service: %v", err)
	}
	pb.RegisterIngesterServer(s, ingester)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
