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
	"strconv"
)

const (
	confluenceContentRetrievalPath = "/wiki/rest/api/content"
	limit                          = 25 // Maximum number of Confluence records to return per page.
	protocol                       = "https"
	service                        = "confluence"
	timestampFormat                = "2006-01-02T15:04:05.000Z"
)

type Contents struct {
	Results []Content
}

type Content struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  struct {
		Storage struct {
			Value string `json:"value"`
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
	stream, err := c.ingesterClient.Ingest(context.Background())
	if err != nil {
		return err
	}

	startIdx := 0
	for {
		endpoint, err := c.constructEndpoint(startIdx)
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

		for _, r := range contents.Results {
			t, err := time.Parse(timestampFormat, r.Version.When)
			modifiedSince, err := ptypes.TimestampProto(t)
			if err != nil {
				return errors.Wrap(err, "failed to parse modified since timestamp as proto")
			}
			content, err := html2text.FromString(r.Body.Storage.Value)
			if err != nil {
				return err
			}
			// TODO: Only retrieve files modified after the last modified time specified in the request.
			doc := pb.Document{
				Id:           &pb.DocumentId{Id: r.Id},
				Title:        r.Title,
				Description:  "",
				Service:      service,
				Content:      content,
				Url:          fmt.Sprintf("%s://%s/wiki%s", protocol, c.hostname, r.Links.WebUI),
				LastModified: modifiedSince,
			}
			if err := stream.Send(&doc); err != nil {
				return err
			}
		}

		if len(contents.Results) != limit {
			break
		}
		startIdx += limit

	}

	if _, err := stream.CloseAndRecv(); err != nil {
		// We expect io.EOF once the stream has closed.
		if err != io.EOF {
			return err
		}
	}

	return nil
}

func (c *ConfluenceService) constructEndpoint(startIdx int) (string, error) {
	uri := protocol + "://" + c.hostname
	endpoint, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", err
	}
	endpoint.Path = confluenceContentRetrievalPath

	data := url.Values{}
	expand := []string{"body.storage", "version"}
	start := strconv.FormatInt(int64(startIdx), 10)
	limit := strconv.FormatInt(limit, 10)
	data.Set("expand", strings.Join(expand, ","))
	data.Set("start", start)
	data.Set("limit", limit)
	endpoint.RawQuery = data.Encode()

	return endpoint.String(), nil
}
