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

var ezoneListCmd = &cobra.Command{
	Use:   "ezone",
	Short: "List deployment zones under an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium to list deployment zones of")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment to list deployment zones of")
			os.Exit(1)
		}

		client := getNewClient()
		var ezones []kld.EZone
		_, err := client.ListEZones(consortiumID, environmentID, &ezones)

		if err != nil {
			fmt.Printf("Failed to list ezones. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(ezones)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var ezoneGetCmd = &cobra.Command{
	Use:   "ezone",
	Short: "Retrieves a ezone details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium that the deployment zone belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment that the deployment zone belongs to")
			os.Exit(1)
		}

		if ezoneID == "" {
			fmt.Println("Missing required parameter: --id for the ezone to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var ezone kld.EZone
		res, err := client.GetEZone(consortiumID, environmentID, ezoneID, &ezone)

		validateGetResponse(res, err, "ezone")
	},
}

var ezoneCreateCmd = &cobra.Command{
	Use:   "ezone",
	Short: "Create a ezone",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumID("ezone")
		validateEnvironmentID("ezone")

		client := getNewClient()
		ezone := kld.NewEZone(name, region, cloud)
		res, err := client.CreateEZone(consortiumID, environmentID, &ezone)

		validateCreationResponse(res, err, "ezone")
	},
}

var ezoneDeleteCmd = &cobra.Command{
	Use:   "ezone",
	Short: "Delete a ezone",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium that the deployment zone belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment that the deployment zone belongs to")
			os.Exit(1)
		}

		if ezoneID == "" {
			fmt.Println("Missing required parameter: --id for the deployment zone to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		res, err := client.DeleteEZone(consortiumID, environmentID, ezoneID)

		validateDeletionResponse(res, err, "ezone")
	},
}

func newEZoneListCmd() *cobra.Command {
	flags := ezoneListCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the deployment zones from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the deployment zones from")

	return ezoneListCmd
}

func newEZoneGetCmd() *cobra.Command {
	flags := ezoneGetCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the deployment zone from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the deployment zone from")
	flags.StringVarP(&ezoneID, "ezone", "n", "", "ID of the deployment zone to retrieve")

	return ezoneGetCmd
}

func newEZoneCreateCmd() *cobra.Command {
	flags := ezoneCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the ezone")
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium this deployment zone is created under")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment this deployment zone is created for")
	flags.StringVarP(&region, "region", "r", "", "Region for the new zone")
	flags.StringVarP(&cloud, "cloud", "C", "", "Cloud for the new zone")

	return ezoneCreateCmd
}

func newEZoneDeleteCmd() *cobra.Command {
	flags := ezoneDeleteCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium this deployment zone is created under")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment this deployment zone is created for")
	flags.StringVarP(&ezoneID, "ezone", "n", "", "ID of the deployment zone to retrieve")

	return ezoneDeleteCmd
}
