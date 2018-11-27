// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/cobra"
)

var appCredsCreateCmd = &cobra.Command{
	Use:   "appcreds",
	Short: "Create application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appcreds")
		validateEnvironmentId("appcreds")
		validateMembershipId("appcreds")
		client := getNewClient()
		appcreds := kld.NewAppCreds(membershipId)
		res, err := client.CreateAppCreds(consortiumId, environmentId, &appcreds)
		validateCreationResponse(res, err, "appcreds")
	},
}

var appCredsGetCmd = &cobra.Command{
	Use:   "appcreds",
	Short: "Get application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appcreds")
		validateEnvironmentId("appcreds")
		validateAppCredsId("appcreds")
		client := getNewClient()
		var appcreds kld.AppCreds
		res, err := client.GetAppCreds(consortiumId, environmentId, appCredsId, &appcreds)
		validateGetResponse(res, err, "appcreds")
	},
}

var appCredsDeleteCmd = &cobra.Command{
	Use:   "appcreds",
	Short: "Delete application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appcreds")
		validateEnvironmentId("appcreds")
		validateAppCredsId("appcreds", false)
		client := getNewClient()
		res, err := client.DeleteAppCreds(consortiumId, environmentId, appCredsId)
		validateDeletionResponse(res, err, "appcreds")
	},
}

var appCredsListCmd = &cobra.Command{
	Use:   "appcreds",
	Short: "List the application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appcreds")
		validateEnvironmentId("appcreds")
		client := getNewClient()
		var appcreds []kld.AppCreds
		_, err := client.ListAppCreds(consortiumId, environmentId, &appcreds)
		if err != nil {
			fmt.Printf("Failed to list app credentials. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(appcreds)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

func newAppCredsCreateCmd() *cobra.Command {
	flags := appCredsCreateCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&membershipId, "membership", "m", "", "Id of the membership to issue the application credentials to")
	return appCredsCreateCmd
}

func newAppCredsGetCmd() *cobra.Command {
	flags := appCredsGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&appCredsId, "appcreds", "a", "", "Id of the API key")
	return appCredsGetCmd
}

func newAppCredsDeleteCmd() *cobra.Command {
	flags := appCredsDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&appCredsId, "appcreds", "a", "", "Id of the API key")
	return appCredsDeleteCmd
}

func newAppCredsListCmd() *cobra.Command {
	flags := appCredsListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	return appCredsListCmd
}
