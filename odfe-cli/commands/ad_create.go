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
	handler "es-cli/odfe-cli/handler/ad"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	commandCreate = "create"
	generate      = "generate-template"
)

//createCmd creates detectors based on file configuration, if interactive mode is on,
//this command will prompt for confirmation on number of detectors will be created on executions.
var createCmd = &cobra.Command{
	Use:   commandCreate + " [detector-configuration-file-path ...] [flags]",
	Short: "Create detectors based on configurations using file",
	Long: fmt.Sprintf("Description:\n  " +
		`Create detectors based on a configuration specified in file path
  To begin, use --generate-template flag to generate sample configuration, save this template locally and update based on your use case.
  use 'odfe-cli ad create file-path' to create detector.`),
	Run: func(cmd *cobra.Command, args []string) {
		generate, _ := cmd.Flags().GetBool(generate)
		if generate {
			generateTemplate()
			return
		}
		//If no args, display usage
		if len(args) < 1 {
			fmt.Println(cmd.Usage())
			return
		}
		err := createDetectors(args)
		DisplayError(err, commandCreate)
	},
}

func generateTemplate() {
	detector, _ := handler.GenerateAnomalyDetector()
	fmt.Println(string(detector))
}

func init() {
	GetADCommand().AddCommand(createCmd)
	createCmd.Flags().BoolP(generate, "g", false, "output sample detector configuration")

}

func createDetectors(fileNames []string) error {

	commandHandler, err := GetADHandler()
	if err != nil {
		return err
	}
	for _, name := range fileNames {
		err = handler.CreateAnomalyDetector(commandHandler, name)
		if err != nil {
			return err
		}
	}
	return nil
}
