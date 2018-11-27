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

var consortiumListCmd = &cobra.Command{
	Use:   "consortium",
	Short: "List the consortiums under the user's account",
	Run: func(cmd *cobra.Command, args []string) {
		client := getNewClient()
		var consortiums []kld.Consortium
		_, err := client.ListConsortium(&consortiums)
		if err != nil {
			fmt.Printf("Failed to list consortiums. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(consortiums)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var consortiumGetCmd = &cobra.Command{
	Use:   "consortium",
	Short: "Get the consortium details",
	Run: func(cmd *cobra.Command, args []string) {
		client := getNewClient()
		var consortium kld.Consortium
		res, err := client.GetConsortium(consortiumId, &consortium)
		validateGetResponse(res, err, "consortium")
	},
}

var consortiumCreateCmd = &cobra.Command{
	Use:   "consortium",
	Short: "Create a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()

		if mode != "single-org" && mode != "multi-org" {
			fmt.Printf("Invalid consortium mode: %s\n", mode)
			os.Exit(1)
		}

		client := getNewClient()
		consortium := kld.NewConsortium(name, desc, mode)
		res, err := client.CreateConsortium(&consortium)
		validateCreationResponse(res, err, "consortium")
	},
}

var consortiumDeleteCmd = &cobra.Command{
	Use:   "consortium",
	Short: "Delete a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateDeleteId("consortium")

		client := getNewClient()
		res, err := client.DeleteConsortium(deleteId)
		validateDeletionResponse(res, err, "consortium")
	},
}

func newConsortiumGetCmd() *cobra.Command {
	flags := consortiumGetCmd.Flags()
	flags.StringVar(&consortiumId, "id", "", "Id of the consortium to retrieve")

	return consortiumGetCmd
}

func newConsortiumListCmd() *cobra.Command {
	return consortiumListCmd
}

func newConsortiumCreateCmd() *cobra.Command {
	flags := consortiumCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the consortium")
	flags.StringVarP(&desc, "desc", "d", "", "Short description of the purpose of the consortium")
	flags.StringVarP(&mode, "mode", "m", "single-org", "single-org or multi-org consortium")

	return consortiumCreateCmd
}

func newConsortiumDeleteCmd() *cobra.Command {
	flags := consortiumDeleteCmd.Flags()
	flags.StringVar(&deleteId, "id", "", "Id of the consortium to delete")

	return consortiumDeleteCmd
}
