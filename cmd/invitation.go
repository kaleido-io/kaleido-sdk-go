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

var invitationListCmd = &cobra.Command{
	Use:   "invitation",
	Short: "List invitations of a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list invitations of")
			os.Exit(1)
		}

		client := getNewClient()
		var invitations []kld.Invitation
		_, err := client.ListInvitations(consortiumId, &invitations)

		if err != nil {
			fmt.Printf("Failed to list invitations. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(invitations)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var invitationGetCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Get invitation details",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("invitation")

		if invitationId == "" {
			fmt.Println("Missing required parameter: --id for the invitation to retrieve")

			os.Exit(1)
		}

		client := getNewClient()
		var invitation kld.Invitation
		res, err := client.GetInvitation(consortiumId, invitationId, &invitation)

		validateGetResponse(res, err, "invitation")
	},
}

var invitationCreateCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Create a invitation for a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateEmail()
		validateConsortiumId("invitation")

		client := getNewClient()
		invitation := kld.NewInvitation(name, email)
		res, err := client.CreateInvitation(consortiumId, &invitation)

		validateCreationResponse(res, err, "invitation")
	},
}

var invitationDeleteCmd = &cobra.Command{
	Use:   "invitation",
	Short: "Delete a invitation from a consortium",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("invitation", false)
		validateDeleteId("invitation")

		client := getNewClient()
		res, err := client.DeleteInvitation(consortiumId, deleteId)

		validateDeletionResponse(res, err, "invitation")
	},
}

func newInvitationListCmd() *cobra.Command {
	flags := invitationListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the invitations from")

	return invitationListCmd
}

func newInvitationGetCmd() *cobra.Command {
	flags := invitationGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the invitation from")
	flags.StringVar(&invitationId, "id", "", "Id of the invitation to retrieve")

	return invitationGetCmd
}

func newInvitationCreateCmd() *cobra.Command {
	flags := invitationCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the organization to create the invitation for")
	flags.StringVarP(&email, "email", "e", "", "Email of the delegate of the organization to create the invitation for")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to add the new invitation to")

	return invitationCreateCmd
}

func newInvitationDeleteCmd() *cobra.Command {
	flags := invitationDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to delete the invitation from")
	flags.StringVar(&deleteId, "id", "", "Id of the invitation to delete")

	return invitationDeleteCmd
}
