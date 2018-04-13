package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
  Use: "get",
  Short: "Get details of a resource: consortium, membership, environment, node",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("get command")
  },
}

func newGetCmd() *cobra.Command {
  getCmd.AddCommand(newConsortiumGetCmd())
  getCmd.AddCommand(newMembershipGetCmd())
  getCmd.AddCommand(newEnvironmentGetCmd())
  getCmd.AddCommand(newNodeGetCmd())

  return getCmd
}