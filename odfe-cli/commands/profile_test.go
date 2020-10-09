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

package commands

import (
	"es-cli/odfe-cli/entity"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func fakeInputProfile(map[string]entity.Profile) entity.Profile {
	return entity.Profile{
		Name:     "default",
		Endpoint: "https://localhost:9200",
		UserName: "admin",
		Password: "admin",
	}
}

func TestCreateProfile(t *testing.T) {
	t.Run("create profile successfully", func(t *testing.T) {
		f, err := ioutil.TempFile("", "profile-create")
		assert.NoError(t, err)
		defer func() {
			err := os.Remove(f.Name())
			assert.NoError(t, err)
		}()
		err = CreateProfile(f.Name(), fakeInputProfile)
		assert.NoError(t, err)
		contents, err := ioutil.ReadFile(f.Name())
		assert.NoError(t, err)
		var config entity.Config
		err = yaml.Unmarshal(contents, &config)
		assert.NoError(t, err)
		assert.EqualValues(t, config, entity.Config{Profiles: []entity.Profile{fakeInputProfile(nil)}})
	})
}
