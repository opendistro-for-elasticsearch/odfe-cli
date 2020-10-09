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
	"es-cli/odfe-cli/controller/config"
	"es-cli/odfe-cli/controller/profile"
	"es-cli/odfe-cli/entity"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

const (
	createNewProfileCommandName = "create"
	profileCommandName          = "profile"
)

//profileCommand is main command for profile operations like list, create and delete
var profileCommand = &cobra.Command{
	Use:   profileCommandName + " sub-command",
	Short: "Manage collection of settings and credentials that you can apply to an odfe-cli command",
	Long: fmt.Sprintf("Description:\n  " +
		`A named profile is a collection of settings and credentials that you can apply to an odfe-cli command.
  When you specify a profile for a command (eg: odfe-cli <command> --profile <profile_name> ), its settings and credentials are used to run that command.
  To configure a default profile for commands, either specify the default profile name in an environment variable (ODFE_PROFILE) or create a profile named 'default'.`),
}

//createProfileCmd creates profile interactively by prompting for name (distinct), user, endpoint, password.
var createProfileCmd = &cobra.Command{
	Use:   createNewProfileCommandName,
	Short: "Creates a new named profile",
	Long:  `Creates a new named profile with details like name, endpoint, user and password`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, err := GetRoot().Flags().GetString(flagConfig)
		if err != nil {
			DisplayError(err, createNewProfileCommandName)
			return
		}
		err = CreateProfile(cfgFile, getNewProfile)
		if err != nil {
			DisplayError(err, createNewProfileCommandName)
			return
		}
	},
}

// init to register commands to its parent command to create a hierarchy
func init() {
	profileCommand.AddCommand(createProfileCmd)
	GetRoot().AddCommand(profileCommand)
}

// CreateProfile creates a new named profile
func CreateProfile(cfgFlagValue string, getNewProfile func(map[string]entity.Profile) entity.Profile) error {
	configFilePath, err := GetConfigFilePath(cfgFlagValue)
	if err != nil {
		return fmt.Errorf("failed to get config file due to: %w", err)
	}
	configController := config.New(configFilePath)
	profileController := profile.New(configController)
	profiles, err := profileController.GetProfilesMap()
	if err != nil {
		return fmt.Errorf("failed to get profile names due to: %w", err)
	}
	newProfile := getNewProfile(profiles)
	if err = profileController.CreateProfile(newProfile); err != nil {
		return fmt.Errorf("failed to create profile %v due to: %w", newProfile, err)
	}
	return nil
}

// getNewProfile gets new profile information from user using command line
func getNewProfile(profileMap map[string]entity.Profile) entity.Profile {
	var name string
	for {
		fmt.Printf("Enter profile's name: ")
		name = getUserInputAsText(checkInputIsNotEmpty)
		if _, ok := profileMap[name]; !ok {
			break
		}
		fmt.Println("profile ", name, "already exists.")
	}
	fmt.Printf("Elasticsearch Endpoint: ")
	endpoint := getUserInputAsText(checkInputIsNotEmpty)
	fmt.Printf("User Name: ")
	user := getUserInputAsText(checkInputIsNotEmpty)
	fmt.Printf("Password: ")
	password := getUserInputAsMaskedText(checkInputIsNotEmpty)
	return entity.Profile{
		Name:     name,
		Endpoint: endpoint,
		UserName: user,
		Password: password,
	}
}

// getUserInputAsText get value from user as text
func getUserInputAsText(isValid func(string) bool) string {
	var response string
	//Ignore return value since validation is applied below
	_, _ = fmt.Scanln(&response)
	if !isValid(response) {
		return getUserInputAsText(isValid)
	}
	return strings.TrimSpace(response)
}

// checkInputIsNotEmpty checks whether input is empty or not
func checkInputIsNotEmpty(input string) bool {
	if len(input) < 1 {
		fmt.Print("value cannot be empty. Please enter non-empty value")
		return false
	}
	return true
}

// getUserInputAsMaskedText get value from user as masked text, since credentials like password
// should not be displayed on console for security reasons
func getUserInputAsMaskedText(isValid func(string) bool) string {
	maskedValue, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	value := fmt.Sprintf("%s", maskedValue)
	if !isValid(value) {
		return getUserInputAsMaskedText(isValid)
	}
	fmt.Println()
	return value
}
