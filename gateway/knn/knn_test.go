package knn

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"odfe-cli/client"
	"odfe-cli/client/mocks"
	"odfe-cli/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestClient(t *testing.T, url string, code int) *client.Client {
	return mocks.NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), url)
		assert.EqualValues(t, len(req.Header), 2)
		return &http.Response{
			StatusCode: code,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString("response")),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
}

func TestGateway_SearchDistinctValues(t *testing.T) {
	ctx := context.Background()
	t.Run("full stats succeeded", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn/stats", 200)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "admin",
			Password: "admin",
		})
		actual, err := testGateway.GetStatistics(ctx, "", "")
		assert.NoError(t, err)
		assert.EqualValues(t, string(actual), "response")
	})
	t.Run("filtered stats succeeded", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn/node1,node2/stats/stat1", 200)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "admin",
			Password: "admin",
		})
		actual, err := testGateway.GetStatistics(ctx, "node1,node2", "stat1")
		assert.NoError(t, err)
		assert.EqualValues(t, string(actual), "response")
	})
	t.Run("gateway failed due to bad user config", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn/stats", 400)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "",
			Password: "admin",
		})
		_, err := testGateway.GetStatistics(ctx, "", "")
		assert.EqualError(t, err, "user name and password cannot be empty")
	})
}
