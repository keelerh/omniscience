package gdrive

import (
	"bufio"
	"context"
	"net/http"
	pb "github.com/keelerh/omniscience/protos"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"strings"
	"fmt"
)

type GoogleDriveService struct {
	svc *drive.Service
}

const GoogleDriveWebViewLink = "https://docs.google.com/document/d/"

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
			// Only attempt to download text files and gdocs
			isGoogleDoc := isGoogleDoc(f.MimeType)
			if !(isTextMime(f.MimeType) || isGoogleDoc) {
				continue
			}
			words, err := downloadFile(g.svc, f.Id, isGoogleDoc)
			if err != nil {
				log.Warningf("Unable to download file: FileId(%v) %v", f.Id, err)
			}
			// TODO: Only retrieve files modified after the last modified time specified in the request.
			doc := pb.Document{
				Id:          f.Id,
				Name:        f.Name,
				Description: f.Description,
				Service:     pb.Service_GDRIVE,
				Url:         GoogleDriveWebViewLink + f.Id,
				Words:       words,
			}
			fmt.Println(f.Id)
			fmt.Println(f.Name)
			fmt.Println(doc)
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

func downloadFile(svc *drive.Service, fileId string, isGoogleDoc bool) ([]string, error) {
	var resp *http.Response
	var err error
	if isGoogleDoc {
		resp, err = svc.Files.Export(fileId, "text/plain").Download()
	} else {
		resp, err = svc.Files.Get(fileId).Download()
	}
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

func isTextMime(mimeType string) bool {
	return strings.HasPrefix("text", mimeType)
}

func isGoogleDoc(mimeType string) bool {
	return mimeType == "application/vnd.google-apps.document"
}
