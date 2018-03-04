package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/keelerh/omniscience/pkg/document_fetcher"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/api/drive/v3"
)

const (
	address                       = "localhost:50051"
	defaultModifiedSinceTimestamp = "02 Jan 06 15:04 MST"
)

var (
	fGoogleServiceAccountPath = flag.String(
		"google_service_account_path",
		"/Users/keeley/google_service_account.json",
		"The path to the Google Drive service account JSON file.")
)

func main() {
	timestamp := defaultModifiedSinceTimestamp
	if len(os.Args) > 1 {
		timestamp = os.Args[1]
	}
	modifiedSince, err := time.Parse(time.RFC822, timestamp)
	if err != nil {
		log.Fatalf("unable to parse modified since timestamp: %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	cfg, err := readGoogleServiceAccountCfg()
	if err != nil {
		log.Fatalf("unable to read Google service account configuration: %v", err)
	}

	gDriveSvc := document_fetcher.NewGoogleDrive(cfg)
	if err = gDriveSvc.Fetch(modifiedSince); err != nil {
		log.Fatalf("failed to fetch documents for Google Drive: %v", err)
	}
}

func readGoogleServiceAccountCfg() (*jwt.Config, error) {
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
