package cmd

import (
	"fmt"
	"os"
	"encoding/json"

	kld "github.com//kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/cobra"
)

var appKeyCreateCmd = &cobra.Command{
	Use:   "appKey",
	Short: "Create application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appKey")
		validateEnvironmentId("appKey")
		validateMembershipId("appKey")
		client := getNewClient()
		appKey := kld.NewAppKey(membershipId)
		res, err := client.CreateAppKey(consortiumId, environmentId, &appKey)
		validateCreationResponse(res, err, "appKey")
	},
}

var appKeyGetCmd = &cobra.Command{
	Use: "appKey",
	Short: "Get application credentials",
	Run: func(cmd *cobra.Command, args []string){
		validateConsortiumId("appKey")
		validateEnvironmentId("appKey")
		validateApiKeyId("appKey")
		client := getNewClient()
		var appKey kld.AppKey
		res, err := client.GetAppKey(consortiumId, environmentId, apiKeyId, &appKey)
		validateGetResponse(res, err, "appKey")
	},
}

var appKeyDeleteCmd = &cobra.Command{
	Use:   "appKey",
	Short: "Delete application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appKey")
		validateEnvironmentId("appKey")
		validateApiKeyId("appKey", false)
		client := getNewClient()
		res, err := client.DeleteAppKey(consortiumId, environmentId, apiKeyId)
		validateDeletionResponse(res, err, "appKey")
	},
}

var appKeyListCmd = &cobra.Command{
	Use: "appKey",
	Short: "List the application credentials",
	Run: func(cmd *cobra.Command, args []string) {
		validateConsortiumId("appKey")
		validateEnvironmentId("appKey")
		client := getNewClient()
		var appKeys []kld.AppKey
		_, err := client.ListAppKeys(consortiumId, environmentId, &appKeys)
		if err != nil {
      fmt.Printf("Failed to list app credentials. %v\n", err)
      os.Exit(1)
    }

    encoded, _ := json.Marshal(appKeys)
    fmt.Printf("\n%+v\n", string(encoded))
	},
}

func newAppKeyCreateCmd() *cobra.Command {
	flags := appKeyCreateCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&membershipId, "membership", "m", "", "Id of the membership this node belongs to")
	return appKeyCreateCmd
}

func newAppKeyGetCmd() *cobra.Command {
	flags := appKeyGetCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&apiKeyId, "apiKey", "a", "", "Id of the API key")
	return appKeyGetCmd
}

func newAppKeyDeleteCmd() *cobra.Command {
	flags := appKeyDeleteCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	flags.StringVarP(&apiKeyId, "apiKey", "a", "", "Id of the API key")
	return appKeyDeleteCmd
}


func newAppKeyListCmd() *cobra.Command {
	flags := appKeyListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to to add the application credentials")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to add the application credentials")
	return appKeyListCmd
}

