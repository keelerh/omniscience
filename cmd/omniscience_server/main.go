// protoc -I protos/ protos/*.proto --go_out=plugins=grpc:protos

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/keelerh/omniscience/pkg/document_services"
	"github.com/keelerh/omniscience/pkg/ingestion"
	pb "github.com/keelerh/omniscience/protos"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/drive/v3"
	"google.golang.org/grpc"
	"github.com/olivere/elastic"
)

const (
	port = ":50051"
)

var (
	fGoogleServiceAccountPath = flag.String(
		"google_service_account_path",
		"/Users/keeley/google_service_account.json",
		"The path to the Google Drive service account JSON file.")
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	cfg, err := readGDriveServiceAccountCfg()
	if err != nil {
		log.Fatalf("unable to read Google Drive service account configuration: %v", err)
	}

	gDriveSvc := document_services.NewGoogleDrive(cfg)
	if err != nil {
		log.Fatalf("failed to instantiate Google Drive service: %v", err)
	}
	pb.RegisterGoogleDriveServer(s, gDriveSvc)

	// Obtain a client and connect to the default Elasticsearch installation on 127.0.0.1:9200.
	esClient, err := elastic.NewClient()
	if err != nil {
		log.Fatalf("failed to initialise Elasticsearch client: %v", err)
	}
	ingestionService, err := ingestion.NewIngestionService(esClient)
	if err != nil {
		log.Fatalf("failed to instantiate Ingestion service: %v", err)
	}
	pb.RegisterIngestionServer(s, ingestionService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func readGDriveServiceAccountCfg() (*jwt.Config, error) {
	dat, err := ioutil.ReadFile(*fGoogleServiceAccountPath)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(dat, drive.DriveReadonlyScope)
	if err != nil {
		return nil, err
	}

	return config, nil
}
