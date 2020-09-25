/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package entity

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const testFileName = "config.yaml"
const testFolderName = "testdata"

func TestConfigFileDeserialization(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		path := filepath.Join(testFolderName, testFileName)
		contents, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var config Config
		err = yaml.Unmarshal(contents, &config)
		assert.NoError(t, err)
		expected := Config{
			Profiles: []Profile{
				{
					Name:     "local",
					Endpoint: "https://localhost:9200",
					UserName: "admin", Password: "admin",
				},
				{
					Name:     "default",
					Endpoint: "https://127.0.0.1:9200",
					UserName: "dadmin", Password: "dadmin",
				},
			}}
		assert.EqualValues(t, expected, config)
	})
}
