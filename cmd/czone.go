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

var czoneListCmd = &cobra.Command{
	Use:   "czone",
	Short: "List deployment zones under an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list deployment zones of")
			os.Exit(1)
		}

		client := getNewClient()
		var czones []kld.CZone
		_, err := client.ListCZones(consortiumId, &czones)

		if err != nil {
			fmt.Printf("Failed to list czones. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(czones)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var czoneGetCmd = &cobra.Command{
	Use:   "czone",
	Short: "Retrieves a czone details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the deployment zone belongs to")
			os.Exit(1)
		}

		if czoneId == "" {
			fmt.Println("Missing required parameter: --id for the czone to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var czone kld.CZone
		res, err := client.GetCZone(consortiumId, czoneId, &czone)

		validateGetResponse(res, err, "czone")
	},
}

var czoneCreateCmd = &cobra.Command{
	Use:   "czone",
	Short: "Create a czone",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("czone")

		client := getNewClient()
		czone := kld.NewCZone(name, region, cloud)
		res, err := client.CreateCZone(consortiumId, &czone)

		validateCreationResponse(res, err, "czone")
	},
}

var czoneDeleteCmd = &cobra.Command{
	Use:   "czone",
	Short: "Delete a czone",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the deployment zone belongs to")
			os.Exit(1)
		}

		if czoneId == "" {
			fmt.Println("Missing required parameter: --id for the deployment zone to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		res, err := client.DeleteCZone(consortiumId, czoneId)

		validateDeletionResponse(res, err, "czone")
	},
}

func newCZoneListCmd() *cobra.Command {
	flags := czoneListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the deployment zones from")

	return czoneListCmd
}

func newCZoneGetCmd() *cobra.Command {
	flags := czoneGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the deployment zone from")
	flags.StringVarP(&czoneId, "czone", "n", "", "Id of the deployment zone to retrieve")

	return czoneGetCmd
}

func newCZoneCreateCmd() *cobra.Command {
	flags := czoneCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the czone")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this deployment zone is created under")
	flags.StringVarP(&region, "region", "r", "", "Region for the new zone")
	flags.StringVarP(&cloud, "cloud", "C", "", "Cloud for the new zone")

	return czoneCreateCmd
}

func newCZoneDeleteCmd() *cobra.Command {
	flags := czoneDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this deployment zone is created under")
	flags.StringVarP(&czoneId, "czone", "n", "", "Id of the deployment zone to retrieve")

	return czoneDeleteCmd
}
