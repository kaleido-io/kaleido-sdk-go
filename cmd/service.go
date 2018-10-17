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

var serviceListCmd = &cobra.Command{
	Use:   "services",
	Short: "List deployed services in an environment",
	Run: func(cmd *cobra.Command, args []string) {
		client := getNewClient()
		var services []kld.Service
		_, err := client.ListServices(consortiumId, environmentId, &services)

		if err != nil {
			fmt.Printf("Failed to list services. %v\n", err)
			os.Exit(1)
		}

		encoded, _ := json.Marshal(services)
		fmt.Printf("\n%+v\n", string(encoded))
	},
}

var serviceGetCmd = &cobra.Command{
	Use:   "service",
	Short: "Retrieves a service details",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getNewClient()
		var services []kld.Service
		res, err := client.GetService(consortiumId, environmentId, serviceId, &services)

		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		return printGetResponse(res, err, "service")
	},
}

var serviceCreateCmd = &cobra.Command{
	Use:   "service",
	Short: "Deploy a service",
	RunE: func(cmd *cobra.Command, args []string) error {
		validateName()
		validateConsortiumId("service")
		validateEnvironmentId("service")
		validateMembershipId("service")

		client := getNewClient()
		service := kld.NewService(service, name, membershipId)
		res, err := client.DeployService(consortiumId, environmentId, &service)

		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		return printCreationResponse(res, err, "service")
	},
}

func newServicesListCmd() *cobra.Command {
	flags := serviceListCmd.Flags()
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium to retrieve the nodes from")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment to retrieve the nodes from")

	serviceListCmd.MarkFlagRequired("consortium")
	serviceListCmd.MarkFlagRequired("environment")

	return serviceListCmd
}

func newServiceGetCmd() *cobra.Command {
	flags := serviceGetCmd.Flags()
	flags.StringVarP(&serviceId, "service-id", "s", "", "Id of the service to retrieve")

	serviceGetCmd.MarkFlagRequired("service-id")

	return serviceGetCmd
}

func newServiceCreateCmd() *cobra.Command {
	flags := serviceCreateCmd.Flags()
	flags.StringVarP(&service, "service", "s", "", "service to deploy (eg: idregistry, hdwallet, ipfs etc.")
	flags.StringVarP(&name, "name", "n", "", "name of the service")
	flags.StringVarP(&membershipId, "membership", "m", "", "Id of the membership this node belongs to")
	flags.StringVarP(&consortiumId, "consortium", "c", "", "Id of the consortium this node is created under")
	flags.StringVarP(&environmentId, "environment", "e", "", "Id of the environment this node is created for")

	serviceCreateCmd.MarkFlagRequired("consortium")
	serviceCreateCmd.MarkFlagRequired("environment")
	serviceCreateCmd.MarkFlagRequired("membership")
	serviceCreateCmd.MarkFlagRequired("name")
	serviceCreateCmd.MarkFlagRequired("service")

	return serviceCreateCmd
}
