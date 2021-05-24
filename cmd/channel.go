// Copyright 2021 Kaleido

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

var members []string

var channelListCmd = &cobra.Command{
	Use:   "channel",
	Short: "List Hyperledger Fabric channels in an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortium for the consortium to list channels of")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environment for the environment to list channels of")
			os.Exit(1)
		}

		client := getNewClient()
		var channels []kld.Channel
		_, err := client.ListChannel(consortiumID, environmentID, &channels)

		if err != nil {
			fmt.Printf("Failed to list channels. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(channels)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var channelGetCmd = &cobra.Command{
	Use:   "channel",
	Short: "Retrieves channel details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortium for the consortium that the channel belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environment for the environment that the channel belongs to")
			os.Exit(1)
		}

		if channelID == "" {
			fmt.Println("Missing required parameter: --id for the channel to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var channel kld.Channel
		res, err := client.GetChannel(consortiumID, environmentID, channelID, &channel)

		validateGetResponse(res, err, "channel")
	},
}

var channelCreateCmd = &cobra.Command{
	Use:   "channel",
	Short: "Create a Hyperledger Fabric channel",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumID("channel")
		validateEnvironmentID("channel")
		validateMembershipID("channel")

		client := getNewClient()
		channel := kld.NewChannel(name, membershipID, members)
		res, err := client.CreateChannel(consortiumID, environmentID, &channel)

		validateCreationResponse(res, err, "channel")
	},
}

func newChannelListCmd() *cobra.Command {
	flags := channelListCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the channels from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the channels from")

	return channelListCmd
}

func newChannelGetCmd() *cobra.Command {
	flags := channelGetCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the channel from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the channel from")
	flags.StringVarP(&channelID, "id", "i", "", "ID of the channel to retrieve")

	return channelGetCmd
}

func newChannelCreateCmd() *cobra.Command {
	flags := channelCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the channel")
	flags.StringVarP(&membershipID, "membership", "m", "", "ID of the membership this channel is initiated by")
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium this channel is created under")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment this channel is created for")
	flags.StringArrayVarP(&members, "members", "", []string{}, "List of membership IDs to be included in the channel")

	return channelCreateCmd
}

func validateMembers() {
	if len(members) == 0 {
		fmt.Println("Missing required parameter: --members for the memberships to add to the new channel")
		os.Exit(1)
	}
}
