package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

type ConsortiumMode int

const (
  SingleOrg   ConsortiumMode = 0
  MultiOrg    ConsortiumMode = 1
)

// for create command
var name string
var desc string
var mode ConsortiumMode

// for delete command
var deleteId string

var consortiumCreateCmd = &cobra.Command{
  Use: "consortium",
  Short: "Create a consortium",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("API URL: %s\n", viper.Get("api.url"))
  },
}

var consortiumDeleteCmd = &cobra.Command{
  Use: "consortium",
  Short: "Delete a consortium",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("API URL: %s\n", viper.Get("api.url"))
  },
}

func newConsortiumCreateCmd() *cobra.Command {
  flags := consortiumCreateCmd.Flags()
  flags.StringVarP(&name, "name", "n", "", "Name of the consortium")
  flags.StringVarP(&desc, "desc", "d", "", "Short description of the purpose of the consortium")

  var modeString string
  flags.StringVarP(&modeString, "mode", "m", "single", "Single-Org (single) or Multi-Org (multi) consortium")
  if modeString == "single" {
    mode = SingleOrg
  } else if modeString == "multi" {
    mode = MultiOrg
  } else {
    panic(fmt.Sprintf("Invalid consortium mode: %n\n", mode))
  }

  return consortiumCreateCmd
}

func newConsortiumDeleteCmd() *cobra.Command {
  flags := consortiumDeleteCmd.Flags()
  flags.StringVar(&deleteId, "id", "", "Id of the consortium to delete")

  return consortiumDeleteCmd
}
