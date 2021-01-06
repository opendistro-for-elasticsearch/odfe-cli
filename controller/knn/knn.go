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
	"context"
	"odfe-cli/gateway/knn"
)

//go:generate go run -mod=mod github.com/golang/mock/mockgen  -destination=mocks/mock_knn.go -package=mocks . Controller

//Controller is an interface for the AD plugin controllers
type Controller interface {
	GetStatistics(context.Context, string, string) ([]byte, error)
}

type controller struct {
	gateway knn.Gateway
}

//GetStatistics gets stats data based on nodes and stat names
func (c controller) GetStatistics(ctx context.Context, nodes string, names string) ([]byte, error) {
	return c.gateway.GetStatistics(ctx, nodes, names)
}

//New returns new Controller instance
func New(gateway knn.Gateway) Controller {
	return &controller{
		gateway,
	}
}
