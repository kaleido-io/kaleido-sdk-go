package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
  Use: "create",
  Short: "Create various resources: consortium, membership, environment, node, appKeys",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("create command")
  },
}

func newCreateCmd() *cobra.Command {
  createCmd.AddCommand(newConsortiumCreateCmd())
  createCmd.AddCommand(newMembershipCreateCmd())
  createCmd.AddCommand(newEnvironmentCreateCmd())
  createCmd.AddCommand(newNodeCreateCmd())
	createCmd.AddCommand(newAppKeyCreateCmd())
	
  return createCmd
}