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

package profile

import (
	"errors"
	config "es-cli/odfe-cli/controller/config/mocks"
	"es-cli/odfe-cli/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func getSampleConfig() entity.Config {
	return entity.Config{
		Profiles: []entity.Profile{{
			Name:     "local",
			Endpoint: "https://localhost:9200",
			UserName: "admin", Password: "admin",
		}}}
}

func TestControllerGetNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(getSampleConfig(), nil)
		ctrl := New(mockConfigCtrl)
		names, err := ctrl.GetProfileNames()
		assert.NoError(t, err)
		assert.EqualValues(t, []string{"local"}, names)
	})
	t.Run("config controller failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(entity.Config{}, errors.New("failed to read"))
		ctrl := New(mockConfigCtrl)
		_, err := ctrl.GetProfileNames()
		assert.EqualError(t, err, "failed to read")
	})
}

func TestControllerGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(getSampleConfig(), nil)
		ctrl := New(mockConfigCtrl)
		p, ok, err := ctrl.GetProfile("local")
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.EqualValues(t, getSampleConfig().Profiles[0], p)
	})
	t.Run("config controller failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(entity.Config{}, errors.New("failed to read"))
		ctrl := New(mockConfigCtrl)
		_, _, err := ctrl.GetProfile("local")
		assert.EqualError(t, err, "failed to read")
	})
	t.Run("no profile found", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(getSampleConfig(), nil)
		ctrl := New(mockConfigCtrl)
		_, ok, err := ctrl.GetProfile("invalid")
		assert.NoError(t, err)
		assert.False(t, ok)
	})
}
func TestControllerCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(entity.Config{}, nil)
		mockConfigCtrl.EXPECT().Write(getSampleConfig()).Return(nil)
		ctrl := New(mockConfigCtrl)
		err := ctrl.CreateProfile(getSampleConfig().Profiles[0])
		assert.NoError(t, err)
	})
	t.Run("config controller read failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(entity.Config{}, errors.New("failed to read"))
		ctrl := New(mockConfigCtrl)
		err := ctrl.CreateProfile(getSampleConfig().Profiles[0])
		assert.EqualError(t, err, "failed to read")
	})
	t.Run("config controller write failed", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockConfigCtrl := config.NewMockController(mockCtrl)
		mockConfigCtrl.EXPECT().Read().Return(entity.Config{}, nil)
		mockConfigCtrl.EXPECT().Write(getSampleConfig()).Return(errors.New("failed to write"))
		ctrl := New(mockConfigCtrl)
		err := ctrl.CreateProfile(getSampleConfig().Profiles[0])
		assert.EqualError(t, err, "failed to write")
	})
}
