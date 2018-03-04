package document_fetcher

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/keelerh/omniscience/protos"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/grpc"
	"google.golang.org/api/drive/v3"
	"sync"
)

const (
	address                = "localhost:50051"
	GoogleDriveWebViewLink = "https://docs.google.com/document/d/"
	service                = "google"
)

type GoogleDriveService struct {
	cfg *jwt.Config
}

func NewGoogleDrive(srvAccountCfg *jwt.Config) *GoogleDriveService {
	return &GoogleDriveService{
		cfg: srvAccountCfg,
	}
}

func (g *GoogleDriveService) Fetch(modifiedSince time.Time) error {
	modifiedSinceProto, err := ptypes.TimestampProto(modifiedSince)
	if err != nil {
		log.Fatalf("unable to parse modified since timestamp as proto: %v", err)
	}

	svc, err := drive.New(g.cfg.Client(context.Background()))
	if err != nil {
		return err
	}

	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer cc.Close()

	client := pb.NewIngesterClient(cc)
	stream, err := client.Ingest(context.Background())

	pageToken := ""
	for {
		q := svc.Files.List()
		// If we have a pageToken set, apply it to the query
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}
		r, err := q.Do()
		if err != nil {
			return err
		}
		var wg sync.WaitGroup
		for _, f := range r.Files {
			// Only attempt to download text files and gdocs
			isGoogleDoc := isGoogleDoc(f.MimeType)
			if !(isTextMime(f.MimeType) || isGoogleDoc) {
				continue
			}
			content, err := downloadFile(svc, f.Id, isGoogleDoc)
			if err != nil {
				log.Warningf("unable to download file: FileId(%v) %v", f.Id, err)
				continue
			}
			// TODO: Only retrieve files modified after the last modified time specified in the request.
			doc := pb.Document{
				Id:          &pb.DocumentId{Id: f.Id},
				Title:       f.Name,
				Description: f.Description,
				Service:     service,
				Content:     content,
				Url:         GoogleDriveWebViewLink + f.Id,
				// TODO: This should be using the ModifiedTIme of the file returned by Google.
				// but that field currently appears to be null.
				LastModified: modifiedSinceProto,
			}
			if err := stream.Send(&doc); err != nil {
				wg.Add(1)
				return err
			}
		}
		pageToken = r.NextPageToken
		if pageToken == "" {
			break
		}
	}

	stream.CloseSend()

	return nil
}

func downloadFile(svc *drive.Service, fileId string, isGoogleDoc bool) (string, error) {
	var resp *http.Response
	var err error
	if isGoogleDoc {
		resp, err = svc.Files.Export(fileId, "text/plain").Download()
	} else {
		resp, err = svc.Files.Get(fileId).Download()
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	content := standardizeSpaces(string(body))

	return content, nil
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func isTextMime(mimeType string) bool {
	return strings.HasPrefix("text", mimeType)
}

func isGoogleDoc(mimeType string) bool {
	return mimeType == "application/vnd.google-apps.document"
}
