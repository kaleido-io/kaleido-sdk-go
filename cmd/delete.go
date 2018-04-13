package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
  Use: "delete",
  Short: "Delete various resources: consortium, membership, environment, node",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Delete command")
  },
}

func newDeleteCmd() *cobra.Command {
  deleteCmd.AddCommand(newConsortiumDeleteCmd())

  return deleteCmd
}