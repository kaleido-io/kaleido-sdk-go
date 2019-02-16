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

var nodeListCmd = &cobra.Command{
	Use:   "node",
	Short: "List nodes under an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium to list nodes of")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --environmentId for the environment to list nodes of")
			os.Exit(1)
		}

		client := getNewClient()
		var nodes []kld.Node
		_, err := client.ListNodes(consortiumId, environmentId, &nodes)

		if err != nil {
			fmt.Printf("Failed to list nodes. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(nodes)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var nodeGetCmd = &cobra.Command{
	Use:   "node",
	Short: "Retrieves a node details",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the node belongs to")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --environmentId for the environment that the node belongs to")
			os.Exit(1)
		}

		if nodeId == "" {
			fmt.Println("Missing required parameter: --id for the node to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var node kld.Node
		res, err := client.GetNode(consortiumId, environmentId, nodeId, &node)

		validateGetResponse(res, err, "node")
	},
}

var nodeCreateCmd = &cobra.Command{
	Use:   "node",
	Short: "Create a node",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumId("node")
		validateEnvironmentId("node")
		validateMembershipId("node")

		client := getNewClient()
		node := kld.NewNode(name, membershipId, ezoneId)
		res, err := client.CreateNode(consortiumId, environmentId, &node)

		validateCreationResponse(res, err, "node")
	},
}

var nodeDeleteCmd = &cobra.Command{
	Use:   "node",
	Short: "Delete a node",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumId == "" {
			fmt.Println("Missing required parameter: --consortiumId for the consortium that the node belongs to")
			os.Exit(1)
		}

		if environmentId == "" {
			fmt.Println("Missing required parameter: --environmentId for the environment that the node belongs to")
			os.Exit(1)
		}

		if nodeId == "" {
			fmt.Println("Missing required parameter: --id for the node to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		res, err := client.DeleteNode(consortiumId, environmentId, nodeId)

		validateDeletionResponse(res, err, "node")
	},
}

func newNodeListCmd() *cobra.Command {
	flags := nodeListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the nodes from")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to retrieve the nodes from")

	return nodeListCmd
}

func newNodeGetCmd() *cobra.Command {
	flags := nodeGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the node from")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to retrieve the node from")
	flags.StringVarP(&nodeId, "node", "n", "", "Id of the node to retrieve")

	return nodeGetCmd
}

func newNodeCreateCmd() *cobra.Command {
	flags := nodeCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the node")
	flags.StringVarP(&membershipId, "membership", "m", "", "Id of the membership this node belongs to")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this node is created under")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment this node is created for")
	flags.StringVarP(&ezoneId, "zone", "z", "", "Id of the environment deployment zone where this node should be created")

	return nodeCreateCmd
}

func newNodeDeleteCmd() *cobra.Command {
	flags := nodeDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this node is created under")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment this node is created for")
	flags.StringVarP(&nodeId, "node", "n", "", "Id of the node to retrieve")

	return nodeDeleteCmd
}
