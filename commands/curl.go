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
	"odfe-cli/client"
	ctrl "odfe-cli/controller/es"
	entity "odfe-cli/entity/es"
	gateway "odfe-cli/gateway/es"
	handler "odfe-cli/handler/es"

	"github.com/spf13/cobra"
)

const (
	curlCommandName         = "curl"
	curlPrettyFlagName      = "pretty"
	curlPathFlagName        = "path"
	curlQueryParamsFlagName = "query-params"
	curlDataFlagName        = "data"
	curlHeadersFlagName     = "headers"
)

//curlCommand is base command for Elasticsearch REST APIs.
var curlCommand = &cobra.Command{
	Use:   curlCommandName,
	Short: "Manage elasticsearch core features",
	Long:  "Use curl command to configure and access features directly",
}

func init() {
	curlCommand.Flags().BoolP("help", "h", false, "Help for curl command")
	curlCommand.PersistentFlags().Bool(curlPrettyFlagName, false, "Response will be formatted")
	GetRoot().AddCommand(curlCommand)
}

//GetCurlCommand returns Curl base command, since this will be needed for subcommands
//to add as parent later
func GetCurlCommand() *cobra.Command {
	return curlCommand
}

//getCurlHandler returns handler by wiring the dependency manually
func getCurlHandler() (*handler.Handler, error) {
	c, err := client.New(nil)
	if err != nil {
		return nil, err
	}
	profile, err := GetProfile()
	if err != nil {
		return nil, err
	}
	g := gateway.New(c, profile)
	facade := ctrl.New(g)
	return handler.New(facade), nil
}

//CurlActionExecute executes API based on user request
func CurlActionExecute(input entity.CurlCommandRequest) error {

	commandHandler, err := getCurlHandler()
	if err != nil {
		return err
	}
	response, err := handler.Curl(commandHandler, input)
	if err == nil {
		fmt.Println(string(response))
		return nil
	}
	if requestError, ok := err.(*entity.RequestError); ok {
		fmt.Println(requestError.GetResponse())
		return nil
	}
	return err
}

func FormatOutput() bool {
	isPretty, _ := curlCommand.PersistentFlags().GetBool(curlPrettyFlagName)
	return isPretty
}

func Run(cmd cobra.Command, cmdName string) {
	input := entity.CurlCommandRequest{
		Action: cmdName,
		Pretty: FormatOutput(),
	}
	input.Path, _ = cmd.Flags().GetString(curlPathFlagName)
	input.QueryParams, _ = cmd.Flags().GetString(curlQueryParamsFlagName)
	input.Data, _ = cmd.Flags().GetString(curlDataFlagName)
	input.Headers, _ = cmd.Flags().GetString(curlHeadersFlagName)
	err := CurlActionExecute(input)
	DisplayError(err, curlGetCommandName)
}
