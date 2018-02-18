package gdrive

import (
	"bufio"
	"context"
	"net/http"
	"strings"

	pb "github.com/keelerh/omniscience/protos"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
)

type GoogleDriveService struct {
	svc *drive.Service
}

func New(client *http.Client) (*GoogleDriveService, error) {
	svc, err := drive.New(client)
	if err != nil {
		return nil, err
	}

	return &GoogleDriveService{
		svc: svc,
	}, nil
}

func (g *GoogleDriveService) GetAll(ctx context.Context, in *pb.GetAllDocumentsRequest) (*pb.GetAllDocumentsResponse, error) {
	var docs []*pb.Document
	pageToken := ""
	for {
		q := g.svc.Files.List()
		// If we have a pageToken set, apply it to the query
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}
		r, err := q.Do()
		if err != nil {
			return &pb.GetAllDocumentsResponse{
				Documents: docs,
			}, err
		}
		for _, f := range r.Files {
			// Only attempt to download text files
			if !strings.HasPrefix(f.MimeType, "text") {
				continue
			}
			words, err := downloadFile(g.svc, f.Id)
			if err != nil {
				log.Warningf("Unable to download file: FileId(%v)", f.Id)
			}
			// TODO: Only retrieve files modified after the last modified time specified in the request.
			doc := pb.Document{
				Id:          f.Id,
				Name:        f.Name,
				Description: f.Description,
				Service:     pb.Service_GDRIVE,
				Url:         f.WebViewLink,
				Words:       words,
			}
			docs = append(docs, &doc)
		}
		pageToken = r.NextPageToken
		if pageToken == "" {
			break
		}
	}

	return &pb.GetAllDocumentsResponse{
		Documents: docs,
	}, nil
}

func downloadFile(svc *drive.Service, fileId string) ([]string, error) {
	resp, err := svc.Files.Get(fileId).Download()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
