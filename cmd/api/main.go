// protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

package main

import (
	"flag"
	"log"
	"net"

	"github.com/keelerh/omniscience/pkg/ingestion"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/olivere/elastic"
	"google.golang.org/grpc"
)

var (
	fElasticsearchUrl = flag.String(
		"elasticsearch_url",
		"http://elasticsearch:9200",
		"Defines the URL endpoint of the Elasticsearch node")
)

// TODO: Port should be set as an ENV variable.
const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Obtain a client and connect to the Elasticsearch installation at the URL endpoint.
	opts := elastic.SetURL(*fElasticsearchUrl)
	esClient, err := elastic.NewClient(opts)
	if err != nil {
		log.Fatalf("failed to initialise Elasticsearch client: %v", err)
	}

	ingester := ingestion.NewIngester(esClient)
	if err != nil {
		log.Fatalf("failed to initialise Ingestion service: %v", err)
	}
	pb.RegisterIngesterServer(s, ingester)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
