package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/keelerh/omniscience/pkg/document_fetcher/github"
	pb "github.com/keelerh/omniscience/protos"
	"google.golang.org/grpc"
)

const (
	address                       = "localhost:50051"
	defaultModifiedSinceTimestamp = "02 Jan 06 15:04 MST"
)

var (
	fGithubApiTokenFilePath = flag.String(
		"github_api_token_file_path",
		"",
		"The path to the Github API token.")
	fGithubOrganization = flag.String(
		"github_organization",
		"",
		"The Github organization from which to index repo READMEs.")
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

	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to IngestionService: %v", err)
	}
	defer cc.Close()

	rawToken, err := ioutil.ReadFile(*fGithubApiTokenFilePath)
	if err != nil {
		log.Fatalf("failed to read Github API token: %v", err)
	}

	token := strings.TrimSpace(string(rawToken))
	client := pb.NewIngesterClient(cc)
	githubSvc := github.NewGithub(token, *fGithubOrganization, &client)

	if err = githubSvc.Fetch(modifiedSince); err != nil {
		log.Fatalf("failed to fetch READMEs from Github: %v", err)
	}
}
