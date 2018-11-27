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

var membershipListCmd = &cobra.Command{
	Use:   "membership",
	Short: "List memberships of a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list memberships of")
			os.Exit(1)
		}

		client := getNewClient()
		var memberships []kld.Membership
		_, err := client.ListMemberships(consortiumId, &memberships)

		if err != nil {
			fmt.Printf("Failed to list memberships. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(memberships)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var membershipGetCmd = &cobra.Command{
	Use:   "membership",
	Short: "Get membership details",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("membership")

		if membershipId == "" {
			fmt.Println("Missing required parameter: --id for the membership to retrieve")

			os.Exit(1)
		}

		client := getNewClient()
		var membership kld.Membership
		res, err := client.GetMembership(consortiumId, membershipId, &membership)

		validateGetResponse(res, err, "membership")
	},
}

var membershipCreateCmd = &cobra.Command{
	Use:   "membership",
	Short: "Create a membership for a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumId("membership")

		client := getNewClient()
		membership := kld.NewMembership(name)
		res, err := client.CreateMembership(consortiumId, &membership)

		validateCreationResponse(res, err, "membership")
	},
}

var membershipDeleteCmd = &cobra.Command{
	Use:   "membership",
	Short: "Delete a membership from a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("membership", false)
		validateDeleteId("membership")

		client := getNewClient()
		res, err := client.DeleteMembership(consortiumId, deleteId)

		validateDeletionResponse(res, err, "membership")
	},
}

func newMembershipListCmd() *cobra.Command {
	flags := membershipListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the memberships from")

	return membershipListCmd
}

func newMembershipGetCmd() *cobra.Command {
	flags := membershipGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the membership from")
	flags.StringVar(&membershipId, "id", "", "Id of the membership to retrieve")

	return membershipGetCmd
}

func newMembershipCreateCmd() *cobra.Command {
	flags := membershipCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the organization to create the membership for")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to add the new membership to")

	return membershipCreateCmd
}

func newMembershipDeleteCmd() *cobra.Command {
	flags := membershipDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to delete the membership from")
	flags.StringVar(&deleteId, "id", "", "Id of the membership to delete")

	return membershipDeleteCmd
}
