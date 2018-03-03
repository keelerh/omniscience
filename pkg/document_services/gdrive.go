package document_services

import (
	"io/ioutil"
	"net/http"
	"strings"

	pb "github.com/keelerh/omniscience/protos"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/drive/v3"
	"golang.org/x/oauth2/jwt"
	"context"
	"fmt"
	"time"
)

type GoogleDriveService struct {
	cfg *jwt.Config
}

const GoogleDriveWebViewLink = "https://docs.google.com/document/d/"

func NewGoogleDrive(srvAccountCfg *jwt.Config) *GoogleDriveService {
	return &GoogleDriveService{
		cfg: srvAccountCfg,
	}
}

func (g *GoogleDriveService) GetAll(request *pb.GetAllDocumentsRequest, stream pb.GoogleDrive_GetAllServer) error {
	ctx, cancelFn := context.WithTimeout(context.Background(), 2 * time.Minute)
	defer cancelFn()

	fmt.Println(g.cfg.Scopes)

	svc, err := drive.New(g.cfg.Client(ctx))
	if err != nil {
		return err
	}

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
				Name:        f.Name,
				Description: f.Description,
				Service:     pb.DocumentService_GDRIVE,
				Content:     content,
				Url:         GoogleDriveWebViewLink + f.Id,
				// TODO: This should be using the ModifiedTIme of the file returned by Google.
				// but that field currently appears to be null.
				LastModified: request.ModifiedSince,
			}
			if err := stream.Send(&doc); err != nil {
				fmt.Println("streaming file...", f.Name)
				return err
			}
		}
		pageToken = r.NextPageToken
		if pageToken == "" {
			break
		}
	}

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
