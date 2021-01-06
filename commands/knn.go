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
	ctrl "odfe-cli/controller/knn"
	gateway "odfe-cli/gateway/knn"
	handler "odfe-cli/handler/knn"

	"github.com/spf13/cobra"
)

const (
	knnCommandName        = "knn"
	knnStatsCommandName   = "stats"
	knnStatsNodesFlagName = "nodes"
	knnStatsNamesFlagName = "stat-names"
)

//knnCommand is base command for k-NN plugin.
var knnCommand = &cobra.Command{
	Use:   knnCommandName,
	Short: "Manage the k-NN plugin short for k-nearest neighbors",
	Long:  "Use the k-NN commands to know about the current status of the KNN Plugin.",
}

//knnStatsCommandName provide stats command for k-NN plugin.
var knnStatsCommand = &cobra.Command{
	Use:   knnStatsCommandName,
	Short: "Manage the k-NN plugin short for k-nearest neighbors",
	Long:  "Use the k-NN commands to know about the current status of the KNN Plugin.",
	Run: func(cmd *cobra.Command, args []string) {
		h, err := GetKNNHandler()
		if err != nil {
			DisplayError(err, knnStatsCommandName)
			return
		}
		nodes, err := cmd.Flags().GetString(knnStatsNodesFlagName)
		if err != nil {
			DisplayError(err, knnStatsCommandName)
			return
		}
		names, err := cmd.Flags().GetString(knnStatsNamesFlagName)
		if err != nil {
			DisplayError(err, knnStatsCommandName)
			return
		}
		err = getStatistics(h, nodes, names)
		DisplayError(err, knnStatsCommandName)
	},
}

func init() {
	//knn base command
	knnCommand.Flags().BoolP("help", "h", false, "Help for k-NN plugin")
	GetRoot().AddCommand(knnCommand)
	//knn stats command
	knnStatsCommand.Flags().BoolP("help", "h", false, "Help for k-NN plugin stats command")
	knnStatsCommand.Flags().StringP(knnStatsNodesFlagName, "n", "", "Input is list of node Ids, separated by ','")
	knnStatsCommand.Flags().StringP(knnStatsNamesFlagName, "s", "", "Input is list of stats names, separated by ','")
	knnCommand.AddCommand(knnStatsCommand)
}

func getStatistics(h *handler.Handler, nodes string, names string) error {
	detector, err := handler.GetStatistics(h, nodes, names)
	if err != nil {
		return nil
	}
	fmt.Println(string(detector))
	return nil
}

//GetKNNHandler returns handler by wiring the dependency manually
func GetKNNHandler() (*handler.Handler, error) {
	c, err := client.New(nil)
	if err != nil {
		return nil, err
	}
	profile, err := GetProfile()
	if err != nil {
		return nil, err
	}
	g := gateway.New(c, profile)
	ctr := ctrl.New(g)
	return handler.New(ctr), nil
}
