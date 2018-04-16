package cmd

import (
  "encoding/json"
  "fmt"
  "os"

  kld "github.com/consensys/photic-sdk-go/kaleido"
  "github.com/spf13/cobra"
)

var nodeListCmd = &cobra.Command{
  Use: "node",
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
  Use: "node",
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
  Use: "node",
  Short: "Create a node",
  Run: func(cmd *cobra.Command, args []string) {
    validateName()
    validateConsortiumId("node")
    validateEnvironmentId("node")
    validateMembershipId("node")

    client := getNewClient()
    node := kld.NewNode(name, membershipId)
    res, err := client.CreateNode(consortiumId, environmentId, &node)

    validateCreationResponse(res, err, "node")
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

  return nodeCreateCmd
}
