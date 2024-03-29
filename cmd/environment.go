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
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium to list environments of")

			os.Exit(1)
		}

		client := getNewClient()
		var environments []kld.Environment
		_, err := client.ListEnvironments(consortiumID, &environments)

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
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium that the environment belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --id for the environment to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var environment kld.Environment
		res, err := client.GetEnvironment(consortiumID, environmentID, &environment)

		validateGetResponse(res, err, "environment")
	},
}

var environmentCreateCmd = &cobra.Command{
	Use:   "environment",
	Short: "Create an environment",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumID("environment")

		client := getNewClient()
		accountBalances := map[string]string{}
		if len(accounts) != len(balances) {
			fmt.Println("For pre-funded accounts, provide accounts (--accounts) and corresponding balances (--balances)")
			os.Exit(1)
		}
		for i := range accounts {
			accountBalances[accounts[i]] = balances[i]
		}
		environment := kld.NewEnvironment(name, desc, provider, consensus, multiRegion, blockPeriod, accountBalances, chainID)
		res, err := client.CreateEnvironment(consortiumID, &environment)

		validateCreationResponse(res, err, "environment")
	},
}

var environmentDeleteCmd = &cobra.Command{
	Use:   "environment",
	Short: "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		validateDeleteID("environment")
		validateConsortiumID("environment", false)

		client := getNewClient()
		res, err := client.DeleteEnvironment(consortiumID, deleteID)

		validateDeletionResponse(res, err, "environment")
	},
}

func newEnvironmentGetCmd() *cobra.Command {
	flags := environmentGetCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the environment from")
	flags.StringVar(&environmentID, "id", "", "ID of the environment to retrieve")

	return environmentGetCmd
}

func newEnvironmentListCmd() *cobra.Command {
	flags := environmentListCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the environments from")

	return environmentListCmd
}

func newEnvironmentCreateCmd() *cobra.Command {
	flags := environmentCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the environment")
	flags.StringVarP(&desc, "desc", "d", "", "Short description of the purpose of the consortium")
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to create the environment under")
	flags.StringVarP(&provider, "provider", "p", "quorum", "underlying protocol to use for this network, quorum or geth")
	flags.StringVarP(&consensus, "consensus", "k", "raft", "consensus algorithm to use for the given protocol, raft or ibft for quorum, poa for geth")
	flags.BoolVarP(&multiRegion, "multi-region", "R", false, "whether to enable multi region")
	flags.IntVarP(&blockPeriod, "block-period", "P", 0, "block period in seconds")
	flags.UintVarP(&chainID, "chain-id", "C", 0, "Chain ID")
	flags.StringArrayVarP(&accounts, "accounts", "a", []string{}, "Account addresses without 0x prefix - for pre-funded accounts")
	flags.StringArrayVarP(&balances, "balances", "b", []string{}, "Account balances for addresses - for pre-funded accounts")

	return environmentCreateCmd
}

func newEnvironmentDeleteCmd() *cobra.Command {
	flags := environmentDeleteCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to delete the environment from")
	flags.StringVar(&deleteID, "id", "", "ID of the environment to delete")

	return environmentDeleteCmd
}
