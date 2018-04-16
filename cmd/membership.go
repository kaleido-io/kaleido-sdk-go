package cmd

import (
  "fmt"
  "os"

  kld "github.com/consensys/photic-sdk-go/kaleido"
  "github.com/spf13/cobra"
)

var membershipGetCmd = &cobra.Command{
  Use: "membership",
  Short: "Get membership details",
  Run: func(cmd *cobra.Command, args []string) {
    validateConsortiumId("membership")

    if membershipId == "" {
      fmt.Println("Missing required parameter: --id for the membership to retrieve")

      os.Exit(1)
    }

    client := getNewClient()
    var membership kld.Membership
    res, err := client.GetMembership(consortiumId, membershipId, &membership)

    validateGetResponse(res, err, "membership")
  },
}

var membershipCreateCmd = &cobra.Command{
  Use: "membership",
  Short: "Create a membership for a consortium",
  Run: func(cmd *cobra.Command, args []string) {
    validateName()
    validateConsortiumId("membership")

    client := getNewClient()
    membership := kld.NewMembership(name)
    res, err := client.CreateMembership(consortiumId, &membership)

    validateCreationResponse(res, err, "membership")
  },
}

var membershipDeleteCmd = &cobra.Command{
  Use: "membership",
  Short: "Delete a membership from a consortium",
  Run: func(cmd *cobra.Command, args []string) {
    validateConsortiumId("membership", false)
    validateDeleteId("membership")

    client := getNewClient()
    res, err := client.DeleteMembership(consortiumId, deleteId)

    validateDeletionResponse(res, err, "membership")
  },
}

func newMembershipGetCmd() *cobra.Command {
  flags := membershipGetCmd.Flags()
  flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the membership from")
  flags.StringVar(&membershipId, "id", "", "Id of the membership to retrieve")

  return membershipGetCmd
}

func newMembershipCreateCmd() *cobra.Command {
  flags := membershipCreateCmd.Flags()
  flags.StringVarP(&name, "name", "n", "", "Name of the organization to create the membership for")
  flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to add the new membership to")

  return membershipCreateCmd
}

func newMembershipDeleteCmd() *cobra.Command {
  flags := membershipDeleteCmd.Flags()
  flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to delete the membership from")
  flags.StringVar(&deleteId, "id", "", "Id of the membership to delete")

  return membershipDeleteCmd
}
