// +build !windows

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

package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	FolderPermission = 0700 // only owner can read, write and execute
	FilePermission   = 0600 // only owner can read and write
)

// createDefaultConfigFolderIfNotExists creates default config file along with folder if
// it doesn't exists
func createDefaultConfigFileIfNotExists() error {
	defaultFilePath := GetDefaultConfigFilePath()
	if isExists(defaultFilePath) {
		return nil
	}
	folderPath := filepath.Dir(defaultFilePath)
	if !isExists(folderPath) {
		err := os.Mkdir(folderPath, FolderPermission)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(defaultFilePath)
	if err != nil {
		return err
	}
	if err = f.Chmod(FilePermission); err != nil {
		return err
	}
	return f.Close()
}

func checkConfigFilePermission(configFilePath string) error {
	//check for config file permission
	info, err := os.Stat(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to get config file info due to: %w", err)
	}
	mode := info.Mode().Perm()

	if mode != FilePermission {
		return fmt.Errorf("permissions %o for '%s' are too open. It is required that your config file is NOT accessible by others", mode, configFilePath)
	}
	return nil
}
