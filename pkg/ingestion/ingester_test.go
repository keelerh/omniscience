package ingestion_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/keelerh/omniscience/pkg/ingestion"
	pb "github.com/keelerh/omniscience/protos"
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/require"
	"github.com/keelerh/omniscience/pkg/ingestion/mocks"
)

func TestIngester_Ingest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	stream := ingester_mocks.NewMockIngester_IngestServer(mockCtrl)
	doc := &pb.Document{
		Id:          &pb.DocumentId{Id: "123"},
		Title:       "Fake document",
		Description: "This is my fake document.",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit," +
			"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." +
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris" +
			"nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in" +
			"reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla" +
			"nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt" +
			"in culpa qui officia deserunt mollit anim id est laborum.",
		Url:          "https://abc.xyz",
		Service:      "fake_service",
		LastModified: ptypes.TimestampNow(),
	}
	stream.EXPECT().Recv().Return(doc, nil).Times(1)
	stream.EXPECT().Recv().Return(nil, io.EOF).Times(1)

	ts := httptest.NewServer(mockHandler())
	defer ts.Close()

	elasticClient, err := mockElasticClient(ts.URL)
	require.NoError(t, err)

	ingester := ingestion.NewIngester(elasticClient)
	err = ingester.Ingest(stream)

	require.NoError(t, err)
}

func mockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := `{
            "took": 73,
            "timed_out": false,
            "hits": [],
            "aggregations": {}
        }`
		w.Write([]byte(resp))
	}
}

func mockElasticClient(url string) (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return client, nil
}
