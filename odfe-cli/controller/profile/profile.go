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
	"es-cli/odfe-cli/controller/config"
	"es-cli/odfe-cli/entity"
)

// go:generate mockgen -destination=mocks/mock_ad.go -package=mocks . Controller
type Controller interface {
	GetProfileNames() ([]string, error)
	GetProfile(name string) (entity.Profile, bool, error)
	CreateProfile(profile entity.Profile) error
}

type controller struct {
	configCtrl config.Controller
}

// New returns new Controller instance
func New(c config.Controller) Controller {
	return &controller{
		configCtrl: c,
	}
}

// GetProfileNames gets list of profile names
func (c controller) GetProfileNames() ([]string, error) {
	data, err := c.configCtrl.Read()
	if err != nil {
		return nil, err
	}
	var names []string
	for _, profile := range data.Profiles {
		names = append(names, profile.Name)
	}
	return names, nil
}

// GetProfile gets the profile named by the name. If the profile is present
// in the config, the profile is returned and the boolean is true.
// Otherwise the returned profile will be empty and the boolean will
// be false.
func (c controller) GetProfile(name string) (entity.Profile, bool, error) {
	data, err := c.configCtrl.Read()
	if err != nil {
		return entity.Profile{}, false, err
	}
	for _, p := range data.Profiles {
		if p.Name == name {
			return p, true, nil
		}
	}
	return entity.Profile{}, false, nil
}

// CreateProfile creates profile and saves it in config file
func (c controller) CreateProfile(p entity.Profile) error {
	data, err := c.configCtrl.Read()
	if err != nil {
		return err
	}
	data.Profiles = append(data.Profiles, p)
	return c.configCtrl.Write(data)
}
