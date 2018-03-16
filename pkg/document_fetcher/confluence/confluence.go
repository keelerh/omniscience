package confluence

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jaytaylor/html2text"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/pkg/errors"
)

const (
	confluenceContentRetrievalPath = "/wiki/rest/api/content"
	protocol                       = "https"
	service                        = "confluence"
)

type Contents struct {
	Results []Content
}

type Content struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Body   struct {
		Storage struct {
			Value          string `json:"value"`
			Representation string `json:"representation"`
		} `json:"storage"`
	} `json:"body"`
	Version struct {
		When string `json:"when"`
	} `json:"version"`
	Links struct {
		WebUI string `json:"webui"`
	} `json:"_links"`
}

type ConfluenceService struct {
	authMethod     AuthMethod
	client         *http.Client
	hostname       string
	ingesterClient pb.IngesterClient
}

func NewConfluence(hostname string, auth AuthMethod, ingesterClient *pb.IngesterClient) (*ConfluenceService, error) {
	return &ConfluenceService{
		authMethod:     auth,
		client:         &http.Client{},
		hostname:       hostname,
		ingesterClient: *ingesterClient,
	}, nil
}

func (c *ConfluenceService) Fetch(modifiedSince time.Time) error {
	modifiedSinceProto, err := ptypes.TimestampProto(modifiedSince)
	if err != nil {
		return errors.Wrap(err, "failed to parse modified since timestamp as proto")
	}

	endpoint, err := c.constructEndpoint()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	res, err := c.sendRequest(req)
	if err != nil {
		return err
	}

	var contents Contents
	if err := json.Unmarshal(res, &contents); err != nil {
		return err
	}
	stream, err := c.ingesterClient.Ingest(context.Background())
	for _, r := range contents.Results {
		//lastModified := r.Version.When
		content, err := html2text.FromString(r.Body.Storage.Value)
		if err != nil {
			return err
		}
		// TODO: Only retrieve files modified after the last modified time specified in the request.
		doc := pb.Document{
			Id:          &pb.DocumentId{Id: r.Id},
			Title:       r.Title,
			Description: "",
			Service:     service,
			Content:     content,
			Url:         fmt.Sprintf("%s://%s/wiki%s", protocol, c.hostname, r.Links.WebUI),
			// TODO: This should be using the When field returned by Confluence.
			LastModified: modifiedSinceProto,
		}
		if err := stream.Send(&doc); err != nil {
			return err
		}
	}

	if _, err := stream.CloseAndRecv(); err != nil {
		// We expect io.EOF once the stream has closed.
		if err != io.EOF {
			return err
		}
	}

	return nil
}

func (c *ConfluenceService) constructEndpoint() (string, error) {
	uri := protocol + "://" + c.hostname
	endpoint, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	endpoint.Path = confluenceContentRetrievalPath
	data := url.Values{}
	expand := []string{"body.storage", "version"}
	data.Set("expand", strings.Join(expand, ","))
	endpoint.RawQuery = data.Encode()

	return endpoint.String(), nil
}
