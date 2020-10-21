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
	"es-cli/odfe-cli/client"
	adctrl "es-cli/odfe-cli/controller/ad"
	esctrl "es-cli/odfe-cli/controller/es"
	adgateway "es-cli/odfe-cli/gateway/ad"
	esgateway "es-cli/odfe-cli/gateway/es"
	handler "es-cli/odfe-cli/handler/ad"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	adCommandName   = "ad"
	flagProfileName = "profile"
)

//adCommand is base command for Anomaly Detection plugin.
var adCommand = &cobra.Command{
	Use:   adCommandName,
	Short: "Manage your Anomaly Detection plugin",
	Long: fmt.Sprintf("Description:\n  " +
		`Use the Anomaly Detectors commands to create, configure and manage Anomaly Detections.
  These commands enables you to leverage Machine Learning based algorithms to automatically detect anomalies as your data is ingested.`),
}

func init() {
	adCommand.PersistentFlags().StringP(flagProfileName, "p", "", "Use a specific profile from your config file.")
	GetRoot().AddCommand(adCommand)
}

//GetADCommand returns AD base command, since this will be needed for subcommands
//to add as parent later
func GetADCommand() *cobra.Command {
	return adCommand
}

//GetADHandler returns handler by wiring the dependency manually
func GetADHandler() (*handler.Handler, error) {
	c, err := client.New(nil)
	if err != nil {
		return nil, err
	}
	p, err := GetProfileController()
	if err != nil {
		return nil, err
	}
	profileFlagValue, err := GetADCommand().Flags().GetString(flagProfileName)
	if err != nil {
		return nil, err
	}
	profile, ok, err := p.GetProfileForExecution(profileFlagValue)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("no profile found for execution. Try %s %s --help for more information", RootCommandName, ProfileCommandName)
	}
	g := adgateway.New(c, &profile)
	esg := esgateway.New(c, &profile)
	esc := esctrl.New(esg)
	ctr := adctrl.New(os.Stdin, esc, g)
	return handler.New(ctr), nil
}
