package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/keelerh/omniscience/pkg/document_fetcher/confluence"
	pb "github.com/keelerh/omniscience/protos"
	"google.golang.org/grpc"
	"io/ioutil"
	"strings"
)

const (
	address                       = "localhost:50051"
	defaultModifiedSinceTimestamp = "02 Jan 06 15:04 MST"
)

var (
	fConfluenceHostname = flag.String(
		"confluence_domain",
		"",
		"A Confluence domain, e.g. your-domain.atlassian.net.")
	fConfluenceUsername = flag.String(
		"confluence_username",
		"",
		"The username associated with the generated Confluence API token.")
	fConfluenceApiTokenFilePath = flag.String(
		"confluence_api_token_file_path",
		"",
		"A Confluence API token; this can be generated at https://id.atlassian.com.")
)

func main() {
	timestamp := defaultModifiedSinceTimestamp
	if len(os.Args) > 1 {
		timestamp = os.Args[1]
	}
	modifiedSince, err := time.Parse(time.RFC822, timestamp)
	if err != nil {
		log.Fatalf("failed to parse modified since timestamp: %v", err)
	}

	rawToken, err := ioutil.ReadFile(*fConfluenceApiTokenFilePath)
	if err != nil {
		log.Fatalf("failed to read Confluence API token: %v", err)
	}
	token := strings.TrimSpace(string(rawToken))

	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to IngestionService: %v", err)
	}
	defer cc.Close()

	client := pb.NewIngesterClient(cc)
	basicAuth := confluence.BasicAuth(*fConfluenceUsername, token)
	confluenceSvc, err := confluence.NewConfluence(*fConfluenceHostname, basicAuth, &client)
	if err != nil {
		log.Fatalf("failed to initialise Confluence service: %v", err)
	}

	if err = confluenceSvc.Fetch(modifiedSince); err != nil {
		log.Fatalf("failed to fetch documents from Confluence: %v", err)
	}
}
