package cmd

import (
  "fmt"

  kld "github.com/consensys/photic-sdk-go/kaleido"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

// for create command
var name string
var desc string
var mode string

// for delete command
var deleteId string

var consortiumCreateCmd = &cobra.Command{
  Use: "consortium",
  Short: "Create a consortium",
  Run: func(cmd *cobra.Command, args []string) {
    if mode != "single-org" && mode != "multi-org" {
      panic(fmt.Sprintf("Invalid consortium mode: %n\n", mode))
    }

    client := kld.NewClient(viper.GetString("api.url"), viper.GetString("api.key"))
    consortium := kld.NewConsortium(name, desc, mode)
    res, err := client.CreateConsortium(&consortium)
    if res.StatusCode() != 201 {
      panic(fmt.Sprintf("Could not create consortium status code: %d.", res.StatusCode()))
    }
    if err != nil {
      panic(err)
    }

    fmt.Printf("\n%+v\n", res)
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
  flags.StringVarP(&mode, "mode", "m", "single-org", "single-org or multi-org consortium")

  return consortiumCreateCmd
}

func newConsortiumDeleteCmd() *cobra.Command {
  flags := consortiumDeleteCmd.Flags()
  flags.StringVar(&deleteId, "id", "", "Id of the consortium to delete")

  return consortiumDeleteCmd
}
