package gdrive

import (
	"context"
	"net/http"

	pb "github.com/keelerh/omniscience/protos"
	"google.golang.org/api/drive/v3"
)

type GoogleDriveService struct {
	service *drive.Service
}

func New(client *http.Client) (*GoogleDriveService, error) {
	service, err := drive.New(client)
	if err != nil {
		return nil, err
	}

	return &GoogleDriveService{service}, nil
}

func (g *GoogleDriveService) GetAll(ctx context.Context, in *pb.GetAllDocumentsRequest) (*pb.GetAllDocumentsResponse, error) {
	var docs []*pb.Document
	pageToken := ""
	for {
		q := g.service.Files.List()
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
			doc := pb.Document{
				Id:          f.Id,
				Name:        f.Name,
				Description: f.Description,
				Service:     pb.Service_GDRIVE,
				Url:         f.WebViewLink,
				// TODO: Retrieve actual content of file.
				// TODO: Only retrieve files modified after the last modified time specified in the request.
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
