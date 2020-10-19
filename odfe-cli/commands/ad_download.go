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
	entity "es-cli/odfe-cli/entity/ad"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	downloadDetectorsCommandName = "download"
	outputDirectoryFlagName      = "output"
	jsonFileExtension            = "json"
)

//downloadDetectorsCmd downloads detectors configuration based on id, name or name regex pattern.
//default input is name pattern, one can change this format to be id by passing --id flag
var downloadDetectorsCmd = &cobra.Command{
	Use:   downloadDetectorsCommandName + " detector_name ..." + " [flags] ",
	Short: "Download detectors based on list of id, name or name regex pattern ",
	Long: fmt.Sprintf("Description:\n  " +
		`Download detectors based on list of user input. Use "" to make sure the name does not match with pwd lists'.
  The default input is detector name, use --id flag if input is detector id instead of name`),
	Run: func(cmd *cobra.Command, args []string) {
		//If no args, display usage
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}
		err := printDetectors(WriteInFile, cmd, args)
		if err != nil {
			DisplayError(err, downloadDetectorsCommandName)
		}
	},
}

//WriteInFile writes detector's configuration on file
//file will be created inside current working directory,
//with detector name as file name
func WriteInFile(cmd *cobra.Command, d *entity.DetectorOutput) error {
	output, _ := cmd.Flags().GetString(outputDirectoryFlagName)
	if _, err := os.Stat(output); os.IsNotExist(err) {
		return fmt.Errorf("output directory [%s] does not exists", output)
	}
	filePath := filepath.Join(output, fmt.Sprintf("%s.%s", d.Name, jsonFileExtension))
	if ok := isCreateFileAllowed(filePath); !ok {
		return nil
	}
	f, err := os.Create(filePath)
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return err
	}
	return FPrint(f, d)
}

//isCreateFileAllowed returns true if no file exists, if file already exists,
//confirms with user, whether to overwrite file or not.
func isCreateFileAllowed(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return askForConfirmation(path)
}

//askForConfirmation get user confirmation before overwriting the file
func askForConfirmation(path string) bool {

	fmt.Printf("overwrite %s? (y/n)", filepath.Base(path))
	var response string
	_, err := fmt.Fscanln(os.Stdin, &response)
	if err != nil {
		//Exit if for some reason, we are not able to accept user input
		fmt.Println("failed to accept value from user due to", err)
		return askForConfirmation(path)
	}
	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no", "\n":
		return false
	default:
		fmt.Println("please enter any of the following, y, yes, n, no.")
		return askForConfirmation(path)
	}
}

func init() {
	GetADCommand().AddCommand(downloadDetectorsCmd)
	downloadDetectorsCmd.Flags().BoolP(deleteDetectorIDFlagName, "", false, "input is detector's id")
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to find current working directory due to ", err)
	}
	downloadDetectorsCmd.Flags().StringP(outputDirectoryFlagName, "o", cwd, "downloads detectors inside this folder path")
}
