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
	DeleteProfileCommandName    = "delete"
	ProfileCommandName          = "profile"
)

//profileCommand is main command for profile operations like list, create and delete
var profileCommand = &cobra.Command{
	Use:   ProfileCommandName + " sub-command",
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
	Long: fmt.Sprintf("Description:\n  " +
		`Creates a new named profile with details like name, endpoint, user and password`),
	Run: func(cmd *cobra.Command, args []string) {
		profileController, err := getController()
		if err != nil {
			DisplayError(err, createNewProfileCommandName)
			return
		}
		err = CreateProfile(profileController, getNewProfile)
		if err != nil {
			DisplayError(err, createNewProfileCommandName)
			return
		}
	},
}

func getController() (profile.Controller, error) {
	cfgFile, err := GetRoot().Flags().GetString(flagConfig)
	if err != nil {
		return nil, err
	}
	return getProfileController(cfgFile)
}

//deleteProfileCmd deletes profiles by names
var deleteProfileCmd = &cobra.Command{
	Use:   DeleteProfileCommandName + " profile_name ...",
	Short: "Deletes profiles by names",
	Long: fmt.Sprintf("Description:\n  " +
		`Deletes profiles by names if it exists in config file, permanently`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := deleteProfiles(cmd, args); err != nil {
			DisplayError(err, DeleteProfileCommandName)
			return
		}
	},
}

func deleteProfiles(cmd *cobra.Command, profiles []string) error {
	if len(profiles) < 1 {
		fmt.Println(cmd.Usage())
		return nil
	}
	profileController, err := getController()
	if err != nil {
		return err
	}
	return profileController.DeleteProfile(profiles)
}

// init to register commands to its parent command to create a hierarchy
func init() {
	profileCommand.AddCommand(createProfileCmd)
	profileCommand.AddCommand(deleteProfileCmd)
	GetRoot().AddCommand(profileCommand)
}

//getProfileController gets profile controller by wiring config controller with config file
func getProfileController(cfgFlagValue string) (profile.Controller, error) {
	configFilePath, err := GetConfigFilePath(cfgFlagValue)
	if err != nil {
		return nil, fmt.Errorf("failed to get config file due to: %w", err)
	}
	configController := config.New(configFilePath)
	profileController := profile.New(configController)
	return profileController, nil
}

// CreateProfile creates a new named profile
func CreateProfile(profileController profile.Controller, getNewProfile func(map[string]entity.Profile) entity.Profile) error {
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
		name = getUserInputAsText("Enter profile's name", checkInputIsNotEmpty)
		if _, ok := profileMap[name]; !ok {
			break
		}
		fmt.Println("profile", name, "already exists.")
	}
	endpoint := getUserInputAsText("Elasticsearch Endpoint", checkInputIsNotEmpty)
	user := getUserInputAsText("User Name", checkInputIsNotEmpty)
	password := getUserInputAsMaskedText("Password", checkInputIsNotEmpty)
	return entity.Profile{
		Name:     name,
		Endpoint: endpoint,
		UserName: user,
		Password: password,
	}
}

// getUserInputAsText get value from user as text
func getUserInputAsText(message string, isValid func(string) bool) string {
	fmt.Printf("%s: ", message)
	var response string
	//Ignore return value since validation is applied below
	_, _ = fmt.Scanln(&response)
	if !isValid(response) {
		return getUserInputAsText("", isValid)
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
func getUserInputAsMaskedText(message string, isValid func(string) bool) string {
	fmt.Printf("%s: ", message)
	maskedValue, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	value := fmt.Sprintf("%s", maskedValue)
	if !isValid(value) {
		return getUserInputAsMaskedText("", isValid)
	}
	fmt.Println()
	return value
}
