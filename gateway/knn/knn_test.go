/*
 * Copyright 2021 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

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

func TestGatewayGetStatistics(t *testing.T) {
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
	t.Run("filtered node and stats succeeded", func(t *testing.T) {

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
	t.Run("filtered node succeeded", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn/node1,node2/stats/", 200)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "admin",
			Password: "admin",
		})
		actual, err := testGateway.GetStatistics(ctx, "node1,node2", "")
		assert.NoError(t, err)
		assert.EqualValues(t, string(actual), "response")
	})
	t.Run("filtered stats succeeded", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn//stats/stat1,stat2", 200)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "admin",
			Password: "admin",
		})
		actual, err := testGateway.GetStatistics(ctx, "", "stat1,stat2")
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
	t.Run("gateway failed due to gateway user config", func(t *testing.T) {

		testClient := getTestClient(t, "http://localhost:9200/_opendistro/_knn/stats", 400)
		testGateway := New(testClient, &entity.Profile{
			Endpoint: "http://localhost:9200",
			UserName: "admin",
			Password: "admin",
		})
		_, err := testGateway.GetStatistics(ctx, "", "")
		assert.Error(t, err)
	})
}
