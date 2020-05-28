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

var environmentListCmd = &cobra.Command{
	Use:   "environment",
	Short: "List environments under a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list environments of")

			os.Exit(1)
		}

		client := getNewClient()
		var environments []kld.Environment
		_, err := client.ListEnvironments(consortiumId, &environments)

		if err != nil {
			fmt.Printf("Failed to list environments. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(environments)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var environmentGetCmd = &cobra.Command{
	Use:   "environment",
	Short: "Retrieves an environment details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the environment belongs to")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --id for the environment to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var environment kld.Environment
		res, err := client.GetEnvironment(consortiumId, environmentId, &environment)

		validateGetResponse(res, err, "environment")
	},
}

var environmentCreateCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create an environment",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumId("environment")

		client := getNewClient()
		environment := kld.NewEnvironment(name, desc, provider, consensus, multiRegion, blockPeriod)
		res, err := client.CreateEnvironment(consortiumId, &environment)

		validateCreationResponse(res, err, "environment")
	},
}

var environmentDeleteCmd = &cobra.Command{
	Use:   "environment",
	Short: "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		validateDeleteId("environment")
		validateConsortiumId("environment", false)

		client := getNewClient()
		res, err := client.DeleteEnvironment(consortiumId, deleteId)

		validateDeletionResponse(res, err, "environment")
	},
}

func newEnvironmentGetCmd() *cobra.Command {
	flags := environmentGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the environment from")
	flags.StringVar(&environmentId, "id", "", "Id of the environment to retrieve")

	return environmentGetCmd
}

func newEnvironmentListCmd() *cobra.Command {
	flags := environmentListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the environments from")

	return environmentListCmd
}

func newEnvironmentCreateCmd() *cobra.Command {
	flags := environmentCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the consortium")
	flags.StringVarP(&desc, "desc", "d", "", "Short description of the purpose of the consortium")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to create the environment under")
	flags.StringVarP(&provider, "provider", "p", "quorum", "underlying protocol to use for this network, quorum or geth")
	flags.StringVarP(&consensus, "consensus", "k", "raft", "consensus algorithm to use for the given protocol, raft or ibft for quorum, poa for geth")
	flags.BoolVarP(&multiRegion, "multi-region", "R", false, "whether to enable multi region")
	flags.IntVarP(&blockPeriod, "block-period", "P", 10, "block period in seconds")

	return environmentCreateCmd
}

func newEnvironmentDeleteCmd() *cobra.Command {
	flags := environmentDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to delete the environment from")
	flags.StringVar(&deleteId, "id", "", "Id of the environment to delete")

	return environmentDeleteCmd
}
