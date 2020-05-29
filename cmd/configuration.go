// Copyright 2020 Kaleido, a ConsenSys business

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
	"io/ioutil"
	"os"

	kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/cobra"
)

var configCreateCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Create configurations",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumID("configurations")
		validateEnvironmentID("configurations")
		validateMembershipID("configurations")
		client := getNewClient()

		if name == "" {
			fmt.Println("Missing required parameter: --name for the configuration to create")
			os.Exit(1)
		}

		if configType == "" {
			fmt.Println("Missing required parameter: --type for the configuration type to create")
			os.Exit(1)
		}

		if detailsFile == "" {
			fmt.Println("Missing required parameter: --file for the JSON details of the configuration type to create")
			os.Exit(1)
		}

		var details map[string]interface{}
		b, err := ioutil.ReadFile(detailsFile)
		if err != nil {
			fmt.Printf("Unable to open file '%s': %s", detailsFile, err)
			os.Exit(1)
		}
		err = json.Unmarshal(b, &details)
		if err != nil {
			fmt.Printf("Unable to parse JSON in file '%s': %s", detailsFile, err)
			os.Exit(1)
		}

		configurations := kld.NewConfiguration(name, membershipID, configType, details)
		res, err := client.CreateConfiguration(consortiumID, environmentID, &configurations)
		validateCreationResponse(res, err, "configurations")
	},
}

var configGetCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Get configurations",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumID("configurations")
		validateEnvironmentID("configurations")

		if configID == "" {
			fmt.Println("Missing required parameter: --id for the configuration to get")
			os.Exit(1)
		}

		client := getNewClient()
		var configuration kld.Configuration
		res, err := client.GetConfiguration(consortiumID, environmentID, configID, &configuration)
		validateGetResponse(res, err, "configurations")
	},
}

var configDeleteCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Delete configurations",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumID("configurations")
		validateEnvironmentID("configurations")

		if configID == "" {
			fmt.Println("Missing required parameter: --id for the configuration to delete")
			os.Exit(1)
		}

		client := getNewClient()
		res, err := client.DeleteConfiguration(consortiumID, environmentID, configID)
		validateDeletionResponse(res, err, "configurations")
	},
}

var configListCmd = &cobra.Command{
	Use:   "configuration",
	Short: "List the configurations",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumID("configurations")
		validateEnvironmentID("configurations")
		client := getNewClient()
		var configurations []kld.Configuration
		_, err := client.ListConfigurations(consortiumID, environmentID, &configurations)
		if err != nil {
			fmt.Printf("Failed to list configurations. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(configurations)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

func newConfigurationCreateCmd() *cobra.Command {
	flags := configCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the configuration")
	flags.StringVarP(&configType, "type", "t", "", "Type of the configuration")
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to to add the configurations")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to add the configurations")
	flags.StringVarP(&membershipID, "membership", "m", "", "ID of the membership to issue the configurations to")
	flags.StringVarP(&detailsFile, "file", "f", "", "JSON file containing type specific details of the configuration to create")
	return configCreateCmd
}

func newConfigurationGetCmd() *cobra.Command {
	flags := configGetCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to to get the configuration")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to get the configuration")
	flags.StringVarP(&configID, "configuration", "i", "", "ID of the configuration")
	return configGetCmd
}

func newConfigurationDeleteCmd() *cobra.Command {
	flags := configDeleteCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to to delete the configurations")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to delete the configurations")
	flags.StringVarP(&configID, "configuration", "i", "", "ID of the configuration")
	return configDeleteCmd
}

func newConfigurationListCmd() *cobra.Command {
	flags := configListCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to to list the configurations")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to list the configurations")
	return configListCmd
}
