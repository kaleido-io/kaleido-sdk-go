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
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium to list nodes of")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment to list nodes of")
			os.Exit(1)
		}

		client := getNewClient()
		var nodes []kld.Node
		_, err := client.ListNodes(consortiumID, environmentID, &nodes)

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
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium that the node belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment that the node belongs to")
			os.Exit(1)
		}

		if nodeID == "" {
			fmt.Println("Missing required parameter: --id for the node to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		var node kld.Node
		res, err := client.GetNode(consortiumID, environmentID, nodeID, &node)

		validateGetResponse(res, err, "node")
	},
}

var nodeCreateCmd = &cobra.Command{
	Use:   "node",
	Short: "Create a node",
	Run: func(cmd *cobra.Command, args []string) {
		validateName()
		validateConsortiumID("node")
		validateEnvironmentID("node")
		validateMembershipID("node")

		client := getNewClient()
		node := kld.NewNode(name, membershipID, ezoneID)

		node.Size = size
		node.KmsID = kmsID
		node.OpsmetricID = opsmetricID
		node.BackupID = backupID
		node.NetworkingID = networkingID
		node.NodeConfigID = nodeConfigID

		res, err := client.CreateNode(consortiumID, environmentID, &node)

		validateCreationResponse(res, err, "node")
	},
}

var nodeDeleteCmd = &cobra.Command{
	Use:   "node",
	Short: "Delete a node",
	Run: func(cmd *cobra.Command, args []string) {
		if consortiumID == "" {
			fmt.Println("Missing required parameter: --consortiumID for the consortium that the node belongs to")
			os.Exit(1)
		}

		if environmentID == "" {
			fmt.Println("Missing required parameter: --environmentID for the environment that the node belongs to")
			os.Exit(1)
		}

		if nodeID == "" {
			fmt.Println("Missing required parameter: --id for the node to retrieve")
			os.Exit(1)
		}

		client := getNewClient()
		res, err := client.DeleteNode(consortiumID, environmentID, nodeID)

		validateDeletionResponse(res, err, "node")
	},
}

func newNodeListCmd() *cobra.Command {
	flags := nodeListCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the nodes from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the nodes from")

	return nodeListCmd
}

func newNodeGetCmd() *cobra.Command {
	flags := nodeGetCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium to retrieve the node from")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment to retrieve the node from")
	flags.StringVarP(&nodeID, "node", "n", "", "ID of the node to retrieve")

	return nodeGetCmd
}

func newNodeCreateCmd() *cobra.Command {
	flags := nodeCreateCmd.Flags()
	flags.StringVarP(&name, "name", "n", "", "Name of the node")
	flags.StringVarP(&membershipID, "membership", "m", "", "ID of the membership this node belongs to")
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium this node is created under")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment this node is created for")
	flags.StringVarP(&ezoneID, "zone", "z", "", "ID of the environment deployment zone where this node should be created")
	flags.StringVarP(&size, "size", "s", "", "Size for the node")
	flags.StringVarP(&kmsID, "kms-id", "k", "", "KMS config ID to attach to the node")
	flags.StringVarP(&opsmetricID, "opsmetric-id", "o", "", "Opsmertic config ID to attach to the node")
	flags.StringVarP(&backupID, "backup-id", "b", "", "Backup config ID to attach to the node")
	flags.StringVarP(&networkingID, "networking-id", "N", "", "Networking config ID to attach to the node")
	flags.StringVarP(&nodeConfigID, "node-config-id", "C", "", "Node config ID to attach to the node")

	return nodeCreateCmd
}

func newNodeDeleteCmd() *cobra.Command {
	flags := nodeDeleteCmd.Flags()
	flags.StringVarP(&consortiumID, "consortium", "c", "", "ID of the consortium this node is created under")
	flags.StringVarP(&environmentID, "environment", "e", "", "ID of the environment this node is created for")
	flags.StringVarP(&nodeID, "node", "n", "", "ID of the node to retrieve")

	return nodeDeleteCmd
}
