package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
  Use: "list",
  Short: "List various resources this user account owns: consortium, membership, environment, node, appKeys",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("list command")
  },
}

func newListCmd() *cobra.Command {
  listCmd.AddCommand(newConsortiumListCmd())
  listCmd.AddCommand(newEnvironmentListCmd())
  listCmd.AddCommand(newNodeListCmd())
  listCmd.AddCommand(newMembershipListCmd())
	listCmd.AddCommand(newAppKeyListCmd())
	
  return listCmd
}